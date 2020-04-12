// Package constructor provides code generator to generate New... function
// that creates an instance of a given struct.
package constructor

import (
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"strings"
	"text/template"

	"github.com/nchern/go-codegen/pkg/code"
)

const (
	initializerTpl = `
func New{{ title .Name }}({{ join .Fields ", "}}) *{{.Name}} {
	return &{{.Name}}{
		{{- range .Fields}}
		{{.Name}}: {{.LName}},
		{{- end}}
	}
}
`
)

type field struct {
	Name string
	Type string
}

func (f *field) String() string {
	return f.LName() + " " + f.Type
}

func (f *field) LName() string {
	return code.ToPackageVisibleIdentifier(f.Name)
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

// Generator implements structure initializer code generator
type Generator struct {
	src io.Reader
}

// FromReader creates generator from reader
func FromReader(r io.Reader) *Generator {
	return &Generator{src: r}
}

// Generate generates the code
func (g *Generator) Generate(w io.Writer) error {
	src, err := code.ReadAndPreparePartialSource(g.src)
	if err != nil {
		return err
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "stdin", src, parser.ParseComments)
	if err != nil {
		if _, ok := err.(scanner.ErrorList); ok {
			// Just ignoer syntax errors - nothing to generate from broken snippet
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
			Funcs(template.FuncMap{"title": strings.Title}).
			Parse(initializerTpl))

		if err := tpl.Execute(w, st); err != nil {
			return err
		}
	}
	return nil
}
