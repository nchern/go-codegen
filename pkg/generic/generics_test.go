package generic

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"

	. "github.com/nchern/go-codegen/pkg/testutil"
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
		{"FromT0To", "FromUUIDTo", TypeMap{T0: "uuid.UUID"}},
		{"T0To", "BarPtrTo", TypeMap{T0: "*model.Bar"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.given, func(t *testing.T) {
			ident := ast.NewIdent(tt.given)
			tt.underTest.substituteTypeVar(&identText{ident})
			assert.Equal(t, tt.expected, ident.Name)
		})
	}
}

func TestStripTypeVarsDecls(t *testing.T) {
	srcText := Text(
		"package main",
		"type T0 string",
		"func foo() {}",
		"type T1 bool")
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
	srcText := Text(
		"package pkg",
		"func foo() {}")
	expectedText := Text(
		"package new",
		"func foo() {}")

	file := CreateGoFile(srcText)
	defer os.Remove(file.Name())

	actualBuf := bytes.Buffer{}
	err := FromFile(file.Name()).
		WithPackageName("new").
		WithTypeMapping(TypeMap{T0: "string"}).
		Generate(&actualBuf)
	assert.NoError(t, err)

	AssertCodeIsSame(t, expectedText, actualBuf.String())
}

func TestGenerateShouldSubsituteTypeVarsAndProduceCode(t *testing.T) {
	srcText := Text(
		"package pkg",
		"type T0 int",
		"func genericFuncT0T1(a string, b T0) (T1, error) {",
		"	m := map[T0]T1{}",
		"	return m[b], nil",
		"}")

	expectedText := Text(
		"package pkg",
		"func genericFuncStringObjectPtr(a string, b string) (*Object, error) {",
		"	m := map[string]*Object{}",
		"	return m[b], nil",
		"}")

	file := CreateGoFile(srcText)
	defer os.Remove(file.Name())

	actualBuf := bytes.Buffer{}
	err := FromFile(file.Name()).
		WithTypeMapping(TypeMap{T0: "string", T1: "*Object"}).
		Generate(&actualBuf)

	assert.NoError(t, err)

	AssertCodeIsSame(t, expectedText, actualBuf.String())
}

func TestGenerateShouldSubsituteTypeVarsInComments(t *testing.T) {
	var tests = []struct {
		name     string
		given    string
		expected string
	}{
		{"single line comment",
			Text(
				"package pkg",
				"type T0 int",
				"",
				"// FooT0 converts T0 to string",
				"func FooT0(a T0) string {",
				"	return a.String()",
				"}"),
			Text(
				"package pkg",
				"// FooBar converts Bar to string",
				"func FooBar(a Bar) string {",
				"	return a.String()",
				"}")},
		{"multiple single line comments",
			Text(
				"package pkg",
				"type T0 int",
				"",
				"// FooT0 converts T0 to string",
				"// T0 should implement Stringer interface",
				"// Warning: T0 should not be nil",
				"func FooT0(a T0) string {",
				"	return a.String()",
				"}"),
			Text(
				"package pkg",
				"// FooBar converts Bar to string",
				"// Bar should implement Stringer interface",
				"// Warning: Bar should not be nil",
				"func FooBar(a Bar) string {",
				"	return a.String()",
				"}")},
		{"multi-line comments",
			Text(
				"package pkg",
				"type T0 int",
				"",
				"/* FooT0 converts T0 to string",
				"   T0 should implement Stringer interface",
				"   Warning: T0 should not be nil */",
				"func FooT0(a T0) string {",
				"	return a.String()",
				"}"),
			Text(
				"package pkg",
				"/* FooBar converts Bar to string",
				"   Bar should implement Stringer interface",
				"   Warning: Bar should not be nil */",
				"func FooBar(a Bar) string {",
				"	return a.String()",
				"}")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actualBuf := bytes.Buffer{}
			err := FromBytes([]byte(tt.given)).
				WithTypeMapping(TypeMap{T0: "Bar"}).
				Generate(&actualBuf)

			assert.NoError(t, err)

			AssertCodeIsSame(t, tt.expected, actualBuf.String())
		})
	}
}
