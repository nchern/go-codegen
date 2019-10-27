package constructor

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"strings"
	"text/template"
	"unicode"
)

const (
	packageHeader = "package main\n"

	initializerTpl = `func New{{.Name}}({{ join .Fields ", "}}) *{{.Name}} {
	return &{{.Name}}{
		{{- range .Fields}}
		{{.Name}}: {{.LName}},
		{{- end}}
	}
}`
)

// Generator abstracts struct initializer generator
type Generator interface {
	WithPackageName(name string) Generator
	Generate(w io.Writer) error
}

type field struct {
	Name string
	Type string
}

func (f *field) String() string {
	return f.LName() + " " + f.Type
}

func (f *field) LName() string {
	isFirst := true
	return strings.Map(
		func(r rune) rune {
			if isFirst {
				isFirst = false
				return unicode.ToLower(r)
			}
			return r
		},
		f.Name)
}

func join(fields []field, sep string) string {
	res := []string{}
	for _, f := range fields {
		res = append(res, f.String())
	}
	return strings.Join(res, sep)
}

type typeInfo struct {
	Name   string
	Fields []field
}

type structInitGenerator struct {
	outputSrc bool
	pkgName   string
	src       io.Reader
}

// FromReader creates generator from reader
func FromReader(r io.Reader) Generator {
	return &structInitGenerator{src: r}
}

// WithPackageName sets the pkg name
func (g *structInitGenerator) WithPackageName(name string) Generator {
	g.pkgName = name
	return g
}

// Generate generates the code
func (g *structInitGenerator) Generate(w io.Writer) error {

	src, err := g.readAndPrepareSource()
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "stdin", src, parser.ParseComments)
	if err != nil {
		return err
	}

	structs := []typeInfo{}

	ast.Inspect(node, func(n ast.Node) bool {
		switch typedNode := n.(type) {
		case *ast.TypeSpec:
			switch t := typedNode.Type.(type) {
			case *ast.StructType:
				info := typeInfo{
					Name: typedNode.Name.String(),
				}
				for _, f := range t.Fields.List {
					if f.Names == nil {
						continue
					}
					info.Fields = append(info.Fields, field{
						Name: f.Names[0].String(),
						Type: fmt.Sprintf("%s", f.Type),
					})
				}
				structs = append(structs, info)
			}
		}
		return true
	})

	if err := g.printInputSourceIfRequired(w, src); err != nil {
		return err
	}

	for _, st := range structs {
		tpl := template.Must(template.New("init").
			Funcs(template.FuncMap{"join": join}).
			Parse(initializerTpl))

		if err := tpl.Execute(w, st); err != nil {
			return err
		}
	}
	return nil
}

func (g *structInitGenerator) printInputSourceIfRequired(w io.Writer, src string) error {
	if _, err := io.WriteString(w, strings.TrimPrefix(src, packageHeader)); err != nil {
		return err
	}
	_, err := fmt.Fprintln(w)
	return err
}

func (g *structInitGenerator) readAndPrepareSource() (string, error) {
	srcBytes, err := ioutil.ReadAll(g.src)
	if err != nil {
		return "", err
	}
	src := string(srcBytes)
	if !strings.HasPrefix(src, "package ") {
		src = packageHeader + src
	}
	return src, nil
}
