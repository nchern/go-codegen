package generic

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/nchern/go-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTypeMapShouldRewriteType(t *testing.T) {
	m := TypeMap{
		T0: "string",
		T1: "*Point",
	}
	var tests = []struct {
		given string

		underTest TypeMap

		expected       string
		expectedResult bool
	}{
		{"T0", m, "string", true},
		{"T1", m, "*Point", true},
		{"int", m, "int", false},
		{"T0", TypeMap{T0: "db.Conn"}, "db.Conn", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run("rewrite:"+tt.given, func(t *testing.T) {
			ident := ast.NewIdent(tt.given)
			assert.Equal(t, tt.expectedResult, tt.underTest.rewriteType(ident))
			assert.Equal(t, tt.expected, ident.Name)
		})
	}
}

func TestTypeMapShouldSubstituteTypeVarInIdent(t *testing.T) {
	tests := []struct {
		given     string
		expected  string
		underTest TypeMap
	}{
		{"NewT0FooT1", "NewStringFooPointPtr", TypeMap{T0: "string", T1: "*Point"}},
		{"NewT0List", "NewObjectList", TypeMap{T0: "interface{}"}},
		{"NewT0List", "NewStringSliceList", TypeMap{T0: "[]string"}},
		{"NewT0List", "NewObjectSliceList", TypeMap{T0: "[]interface{}"}},
		{"NewT0List", "NewFooPtrSliceList", TypeMap{T0: "[]*Foo"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.given, func(t *testing.T) {
			ident := ast.NewIdent(tt.given)
			tt.underTest.substituteTypeVarInIdent(ident)
			assert.Equal(t, tt.expected, ident.Name)
		})
	}
}

func TestStripTypeVarsDecls(t *testing.T) {
	srcText := `
	package main
	type T0 string
	func foo() {}
	type T1 bool
	`
	m := TypeMap{
		T0: "string",
		T1: "*Point",
	}
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", srcText, parser.ParseComments)
	assert.NoError(t, err)

	m.stripTypeVarsDecls(node)

	buf := &bytes.Buffer{}
	assert.NoError(t, printer.Fprint(buf, fset, node))
	assert.Equal(t, "package main\n\nfunc foo()\t{}\n", buf.String())
}

func TestGenerateShouldFailOnUnsupportedTypeVar(t *testing.T) {
	err := FromFile("file.go").
		WithTypeMapping(TypeMap{TypeVar("BOO"): "*string"}).
		Generate(os.Stderr)
	assert.Error(t, err)
}

func TestGenerateWithPackageNameShouldRewritePackageName(t *testing.T) {
	srcText := `package pkg
	func foo() {
	}`
	expectedText := `package new
	func foo() {
	}`
	file := testutil.CreateGoFile(srcText)
	defer os.Remove(file.Name())

	actualBuf := bytes.Buffer{}
	err := FromFile(file.Name()).
		WithPackageName("new").
		WithTypeMapping(TypeMap{T0: "string"}).
		Generate(&actualBuf)
	assert.NoError(t, err)

	testutil.AssertCodeIsSame(t, expectedText, actualBuf.String())
}

func TestGenerateShouldSubsituteTypeVarsAndProduceCode(t *testing.T) {
	srcText := `package pkg
	type T0 int
	func genericFuncT0T1(a string, b T0) (T1, error) {
		m := map[T0]T1{}
		return m[b], nil
	}`

	expectedText := `package pkg
	func genericFuncStringObjectPtr(a string, b string) (*Object, error) {
		m := map[string]*Object{}
		return m[b], nil
	}`

	file := testutil.CreateGoFile(srcText)
	defer os.Remove(file.Name())

	actualBuf := bytes.Buffer{}
	err := FromFile(file.Name()).
		WithTypeMapping(TypeMap{T0: "string", T1: "*Object"}).
		Generate(&actualBuf)

	assert.NoError(t, err)

	testutil.AssertCodeIsSame(t, expectedText, actualBuf.String())
}
