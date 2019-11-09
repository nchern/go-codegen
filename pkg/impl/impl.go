package impl

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"strings"
	"text/template"

	"github.com/nchern/go-codegen/pkg/code"
)

const (
	implTpl = `
type {{.StructName}} struct {}

{{- range .MethodSignatures}}

func ({{$.Reciever}} *{{$.StructName}}) {{.}} {
	panic("Not implemented")
}
{{- end}}
`
)

// TODO: set  custom stub struct name
// TODO: generate default returns based on return types

// Generator implements an interfacer implenetation code generator
type Generator struct {
	outputSrc bool
	src       io.Reader
}

// FromReader returns ImplementationGenerator that reads source from provided reader
func FromReader(r io.Reader) *Generator {
	return &Generator{
		src: r,
	}
}

// WithOutputSrc sets the flag outputSrc
func (g *Generator) WithOutputSrc(outputSrc bool) *Generator {
	g.outputSrc = outputSrc
	return g
}

// Generate generates implenetation of a given interface(s)
func (g *Generator) Generate(w io.Writer) error {
	interfaces := []interfaceInfo{}

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
		return err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch typedNode := n.(type) {
		case *ast.TypeSpec:
			switch t := typedNode.Type.(type) {
			case *ast.InterfaceType:
				iface := interfaceInfo{
					Name: typedNode.Name.String(),
				}
				for _, m := range t.Methods.List {
					if m.Names == nil {
						err = errors.New("Unsupported method signature: empty names")
						return false
					}
					iface.MethodSignatures = append(iface.MethodSignatures, src[m.Pos()-1:m.End()-1])
				}

				interfaces = append(interfaces, iface)
			}
		}
		return true
	})
	if err != nil {
		return err
	}

	for _, iface := range interfaces {
		tpl := template.Must(template.New("init").Parse(implTpl))
		if err := tpl.Execute(w, iface); err != nil {
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

type interfaceInfo struct {
	Name             string
	MethodSignatures []string
}

func (i interfaceInfo) Reciever() string {
	return string(i.StructName()[0])
}

func (i interfaceInfo) StructName() string {
	return code.ToPackageVisibleIdentifier(i.Name)
}
