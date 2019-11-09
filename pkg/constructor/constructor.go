package constructor

import (
	"fmt"
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
	initializerTpl = `func New{{.Name}}({{ join .Fields ", "}}) *{{.Name}} {
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
	outputSrc bool
	src       io.Reader
}

// FromReader creates generator from reader
func FromReader(r io.Reader) *Generator {
	return &Generator{src: r}
}

// WithOutputSrc sets the flag outputSrc
func (g *Generator) WithOutputSrc(outputSrc bool) *Generator {
	g.outputSrc = outputSrc
	return g
}

// Generate generates the code
func (g *Generator) Generate(w io.Writer) error {
	src, err := code.ReadAndPreparePartialSource(g.src)
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

func (g *Generator) printInputSourceIfRequired(w io.Writer, src string) error {
	if !g.outputSrc {
		return nil
	}
	if _, err := io.WriteString(w, strings.TrimPrefix(src, code.PackageMain)); err != nil {
		return err
	}
	_, err := fmt.Fprintln(w)
	return err
}
