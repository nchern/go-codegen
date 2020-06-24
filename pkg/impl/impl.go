// Package impl provides code generator to generate interface implementations
package impl

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/scanner"
	"go/token"
	"io"
	"strings"
	"text/template"

	"github.com/nchern/go-codegen/pkg/code"
)

const (
	implTpl = `
type {{.StructName}} struct {}

{{- range .Methods}}

{{ .Comments }}
func ({{$.Receiver}} *{{$.StructName}}) {{.Signature}} {
	{{ .Body }}
}
{{- end}}
`
)

// TODO: set custom stub struct name
// TODO: generate default returns based on return types

// Generator generates interface implementation code.
type Generator struct {
	src io.Reader
}

// FromReader returns a new instance of Generator that reads source from provided reader
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
					mInfo := methodInfo{
						Signature: src[m.Pos()-1 : m.End()-1],
						Type:      m.Type.(*ast.FuncType),
					}
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

	Type *ast.FuncType
}

func (i methodInfo) Body() string {
	var results *ast.FieldList
	if i.Type != nil {
		results = i.Type.Results
	}
	if results != nil {
		res := results.List[0]

		var buf bytes.Buffer
		fset := token.NewFileSet()

		if err := printer.Fprint(&buf, fset, res.Type); err != nil {
			return fmt.Sprintf("// error during generation: %s", err)
		}

		tp := buf.String()
		switch {
		case tp == code.StringDecl:
			return "return \"\""
		case tp == code.BoolDecl:
			return "return false"
		case tp == code.ObjectDecl:
			return "return nil"
		case strings.HasPrefix(tp, code.IntDecl):
			return "return 0"
		}

		if strings.HasPrefix(tp, code.FloatDecl) {
			return "return 0"
		}

		if strings.HasPrefix(tp, code.PtrDecl) ||
			strings.HasPrefix(tp, code.SliceDecl) {
			return "return nil"
		}

	}
	return "panic(\"Not implemented\")"
}

type interfaceInfo struct {
	Name    string
	Methods []methodInfo
}

func (i interfaceInfo) Receiver() string {
	return string(i.StructName()[0])
}

func (i interfaceInfo) StructName() string {
	return code.ToPackageVisibleIdentifier(i.Name)
}
