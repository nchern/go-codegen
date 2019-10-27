package constructor

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
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
}
`
)

// Generator abstracts struct initializer generator
type Generator interface {
	WithPackageName(name string) Generator
	Generate(w io.Writer) error
	WithOutputSrc(bool) Generator
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

func typeToString(expr ast.Expr, src string) string {
	return src[expr.Pos()-1 : expr.End()-1]
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

// WithOutputSrc sets the flag outputSrc
func (g *structInitGenerator) WithOutputSrc(outputSrc bool) Generator {
	g.outputSrc = outputSrc
	return g
}

// Generate generates the code
func (g *structInitGenerator) Generate(w io.Writer) error {
	src, err := g.readAndPrepareSource()
	if err != nil {
		return err
	}
	if err := g.printInputSourceIfRequired(w, src); err != nil {
		return err
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "stdin", src, parser.ParseComments)
	if err != nil {
		if _, ok := err.(scanner.ErrorList); ok {
			// TODO: decide how to handle it without just skipping
			// for _, v := range syntaxErr {
			//	log.Printf("[%s] %T", v, v)
			// }

			// just do nothing if we can not parse the input - no generation will happen
			return nil
		}
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
						Type: typeToString(f.Type, src),
					})
				}
				structs = append(structs, info)
			}
		}
		return true
	})

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
	if !g.outputSrc {
		return nil
	}
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
