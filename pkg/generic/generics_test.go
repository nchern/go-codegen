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

func TestTypeMap(t *testing.T) {
	m := TypeMap{
		T0: "string",
		T1: "*Point",
	}

	ident := ast.NewIdent("T0")
	assert.True(t, m.rewriteType(ident))
	assert.Equal(t, "string", ident.String())

	ident = ast.NewIdent("int")
	assert.False(t, m.rewriteType(ident))
	assert.Equal(t, "int", ident.String())

	ident = ast.NewIdent("NewT0FooT1")
	m.substituteTypeVarInIdent(ident)
	assert.Equal(t, "NewStringFooPointPtr", ident.Name)

	// test complex cases

	m = TypeMap{T0: "interface{}"}
	ident = ast.NewIdent("NewT0List")
	m.substituteTypeVarInIdent(ident)
	assert.Equal(t, "NewObjectList", ident.Name)

	m = TypeMap{T0: "[]string"}
	ident = ast.NewIdent("NewT0List")
	m.substituteTypeVarInIdent(ident)
	assert.Equal(t, "NewStringSliceList", ident.Name)

	m = TypeMap{T0: "[]interface{}"}
	ident = ast.NewIdent("NewT0List")
	m.substituteTypeVarInIdent(ident)
	assert.Equal(t, "NewObjectSliceList", ident.Name)

	m = TypeMap{T0: "[]*Foo"}
	ident = ast.NewIdent("NewT0List")
	m.substituteTypeVarInIdent(ident)
	assert.Equal(t, "NewFooPtrSliceList", ident.Name)
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

func TestGenericBadTypeVarError(t *testing.T) {
	err := FromFile("file.go").
		WithTypeMapping(TypeMap{TypeVar("BOO"): "*string"}).
		Generate(os.Stderr)
	assert.Error(t, err)
}

func TestGenericWriteToRewritesPackageName(t *testing.T) {
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

	assert.Equal(t, testutil.FormatSrc(expectedText), testutil.FormatSrc(actualBuf.String()))
}

func TestGenericWriteTo(t *testing.T) {
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

	assert.Equal(t, testutil.FormatSrc(expectedText), testutil.FormatSrc(actualBuf.String()))
}
