package impl

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"text/template"

	"github.com/nchern/go-codegen/pkg/code"
)

const (
	implTpl = `
type {{.StructName}} struct {}

{{- range .Methods}}

{{ .Comments }}
func ({{$.Reciever}} *{{$.StructName}}) {{.Signature}} {
	panic("Not implemented")
}
{{- end}}
`
)

// TODO: set custom stub struct name
// TODO: generate default returns based on return types

// Generator implements an interface implenetation code generator
type Generator struct {
	src io.Reader
}

// FromReader returns ImplementationGenerator that reads source from provided reader
func FromReader(r io.Reader) *Generator {
	return &Generator{
		src: r,
	}
}

// Generate generates implenetation of a given interface(s)
func (g *Generator) Generate(w io.Writer) error {
	interfaces := []interfaceInfo{}

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
					mInfo := methodInfo{Signature: src[m.Pos()-1 : m.End()-1]}
					if m.Doc != nil {
						mInfo.Comments = src[m.Doc.Pos()-1 : m.Doc.End()-1]
					}
					iface.Methods = append(iface.Methods, mInfo)
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

type methodInfo struct {
	Comments  string
	Signature string
}

type interfaceInfo struct {
	Name    string
	Methods []methodInfo
}

func (i interfaceInfo) Reciever() string {
	return string(i.StructName()[0])
}

func (i interfaceInfo) StructName() string {
	return code.ToPackageVisibleIdentifier(i.Name)
}
