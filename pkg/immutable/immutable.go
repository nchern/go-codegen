package immutable

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"strings"
	"text/template"
	"unicode"
)

var (
	// ErrBadInputType indicates that given type can not be used to generate immutable structure
	ErrBadInputType = errors.New("Input type must be go interface")

	// ErrUnsupportedMethodSignature indicates that method is not suitble to generate immutable structure
	ErrUnsupportedMethodSignature = errors.New("Method must not be anonymous, have 0 input parameters and 1 output")
)

func camelCaseToSnakeCase(name string) string {
	var i, j int
	var c rune
	tokens := []string{}
	for j, c = range name {
		if j > 0 && unicode.IsUpper(c) {
			tokens = append(tokens, strings.ToLower(name[i:j]))
			i = j
		}
	}
	if i < j || j == i {
		tokens = append(tokens, strings.ToLower(name[i:j+1]))
	}
	return strings.Join(tokens, "_")
}

type method struct {
	Name     string
	RetValue string
}

func (m *method) FieldName() string {
	return m.Name + "Field"
}

func (m *method) StructField() string {
	return fmt.Sprintf("%s %s `json:\"%s\"`", m.FieldName(), m.RetValue, camelCaseToSnakeCase(m.Name))
}

func (m *method) generateImmutableSetter(structName string) string {
	return fmt.Sprintf("func (m *%s) %s() %s { return m.%s }",
		structName, m.Name, m.RetValue, m.FieldName())
}

func (m *method) generateBuilderSetter(builderTypeName string) string {
	immutableFieldName := m.FieldName()
	return fmt.Sprintf("// %s sets corresponding field value\nfunc (b *%s) %s(%s %s) *%s { b.value.%s = %s; return b }",
		m.Name, builderTypeName, m.Name, immutableFieldName, m.RetValue, builderTypeName, immutableFieldName, immutableFieldName)
}

type typeInfo struct {
	Name string

	Methods []method
}

func (t *typeInfo) BuilderName() string {
	return "Immutable" + t.Name + "Builder"
}

func (t *typeInfo) StructName() string {
	return strings.ToLower(t.Name) + "Struct"
}

func (t *typeInfo) generateImmutableStruct(w io.Writer) {
	lines := []string{"type {{.StructName}}  struct {"}
	for _, m := range t.Methods {
		lines = append(lines, m.StructField())
	}
	lines = append(lines, "}")
	for _, m := range t.Methods {
		lines = append(lines, m.generateImmutableSetter(t.StructName()))
	}
	t.generate(lines, w)
}

func (t *typeInfo) generateImmutableBuilder(w io.Writer) {
	lines := []string{
		"type {{.BuilderName}} struct { value *{{.StructName}} }",
		"// New{{.Name}}Builder creates new {{.BuilderName}} builder",
		"func New{{.Name}}Builder() *{{.BuilderName}} { return &{{.BuilderName}}{ &{{.StructName}}{} } }",
	}
	for _, m := range t.Methods {
		lines = append(lines, m.generateBuilderSetter(t.BuilderName()))
	}
	lines = append(lines,
		"// Build builds the immutable structure",
		"func (b {{.BuilderName}}) Build() {{.Name}} { ret := *b.value; return &ret }")
	t.generate(lines, w)
}

func (t *typeInfo) generate(lines []string, w io.Writer) {
	src := strings.Join(append(lines, "\n"), "\n")
	tpl := template.Must(template.New("immutable").Parse(src))
	if err := tpl.Execute(w, t); err != nil {
		panic(err)
	}
}

// Generator implements immutables code generator
type Generator struct {
	pkgName  string
	filename string
}

// FromFile returns ImmutableGenerator from given file
func FromFile(filename string) *Generator {
	return &Generator{
		filename: filename,
	}
}

// WithPackageName sets the package name
func (g *Generator) WithPackageName(name string) *Generator {
	g.pkgName = name
	return g
}

// Generate generates immutable struct anb builder from a given interface
func (g *Generator) Generate(w io.Writer) error {
	immutables := []typeInfo{}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, g.filename, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch typedNode := n.(type) {
		case *ast.TypeSpec:
			switch t := typedNode.Type.(type) {
			case *ast.InterfaceType:
				immutable := typeInfo{
					Name: typedNode.Name.String(),
				}
				for _, m := range t.Methods.List {
					if m.Names == nil {
						err = ErrUnsupportedMethodSignature
						return false
					}
					methodName := m.Names[0].String()
					methodFunc := m.Type.(*ast.FuncType)
					if methodFunc.Params.NumFields() != 0 {
						err = ErrUnsupportedMethodSignature
						return false
					}
					if methodFunc.Results.NumFields() != 1 {
						err = ErrUnsupportedMethodSignature
						return false
					}
					immutable.Methods = append(immutable.Methods, method{
						Name:     methodName,
						RetValue: (methodFunc.Results.List[0].Type.(*ast.Ident)).String(),
					})
				}

				immutables = append(immutables, immutable)
			}
		}
		return true
	})
	if err != nil {
		return err
	}
	if g.pkgName != "" {
		if _, err := io.WriteString(w, fmt.Sprintf("package %s\n", g.pkgName)); err != nil {
			return err
		}
	}
	for _, immutable := range immutables {
		immutable.generateImmutableStruct(w)
		immutable.generateImmutableBuilder(w)
	}

	return nil
}
