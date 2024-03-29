package impl

import (
	"bytes"
	"testing"

	. "github.com/nchern/go-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateImplementation(t *testing.T) {
	var tests = []struct {
		name     string
		given    string
		expected string
	}{
		{
			"with panics as impl",
			Text(
				"type TestInterface interface {",
				"	Foo(u *User)",
				"	Bar(a int, b float64) CustomStruct",
				"}"),
			Text(
				"type testInterface struct {}",
				"func (t *testInterface) Foo(u *User) {",
				"	panic(\"Not implemented\")",
				"}",
				"",
				"func (t *testInterface) Bar(a int, b float64) CustomStruct {",
				"	panic(\"Not implemented\")",
				"}",
			)},
		{
			"return default values for primitive types",
			Text(
				"type TestInterface interface {",
				"	Foo() int",
				"	Bar(a string) bool",
				"	FooBar(a string) string",
				"}"),
			Text(
				"type testInterface struct {}",
				"func (t *testInterface) Foo() int {",
				"	return 0",
				"}",
				"",
				"func (t *testInterface) Bar(a string) bool {",
				"	return false",
				"}",
				"",
				"func (t *testInterface) FooBar(a string) string {",
				"	return \"\"",
				"}",
			)},
		{
			"return default values in case of specific numeric types",
			Text(
				"type TestInterface interface {",
				"	Foo() int32",
				"	Foo64() int64",
				"	Bar(a string) float32",
				"	Bar64(a *testing.T) float64",
				"}"),
			Text(
				"type testInterface struct {}",
				"func (t *testInterface) Foo() int32 {",
				"	return 0",
				"}",
				"",
				"func (t *testInterface) Foo64() int64 {",
				"	return 0",
				"}",
				"",
				"func (t *testInterface) Bar(a string) float32 {",
				"	return 0",
				"}",
				"",
				"func (t *testInterface) Bar64(a *testing.T) float64 {",
				"	return 0",
				"}",
			),
		},
		{
			"return nil as default value for various pointer types",
			Text(
				"type TestInterface interface {",
				"	Foo() interface{}",
				"	Bar(a []string) *CustomStruct",
				"	FooBar(a ...string) []*CustomStruct",
				"}"),
			Text(
				"type testInterface struct {}",
				"func (t *testInterface) Foo() interface{} {",
				"	return nil",
				"}",
				"",
				"func (t *testInterface) Bar(a []string) *CustomStruct {",
				"	return nil",
				"}",
				"",
				"func (t *testInterface) FooBar(a ...string) []*CustomStruct {",
				"	return nil",
				"}",
			)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var actual bytes.Buffer
			assert.NoError(t, FromReader(bytes.NewBufferString(tt.given)).Generate(&actual))

			AssertCodeIsSame(t, tt.expected, actual.String())
		})
	}
}

func TestShouldGenerateNothingOnUnsupportedTypes(t *testing.T) {
	var tests = []struct {
		name  string
		given string
	}{
		{"empty", "\n"},
		{"struct", "type Foo struct {}"},
		{"function", "func foo() error { return nil }"},
		{"just code", "var i = 0\nfmt.Println(i)"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var actual bytes.Buffer
			err := FromReader(bytes.NewBufferString(tt.given)).Generate(&actual)

			assert.NoError(t, err)
			assert.Equal(t, "", actual.String())
		})
	}
}

func TestShouldGenerateMethodsWithCommentsIfCommentsWereProvided(t *testing.T) {
	source := Text(
		"type TestInterface interface {",
		"	// Foo has a single line comment",
		"	Foo() int",

		"	// Bar has",
		"	// two single line comments",
		"	Bar(i int) string",

		"	/*",
		"	Buzz has multi-",
		"	line comment",
		"	*/",
		"	Buzz()",
		"}")

	expected := Text(
		"type testInterface struct {}",

		"// Foo has a single line comment",
		"func (t *testInterface) Foo() int {",
		"	return 0",
		"}",

		"// Bar has",
		"// two single line comments",
		"func (t *testInterface) Bar(i int) string {",
		"	return \"\"",
		"}",

		"/*",
		"Buzz has multi-",
		"line comment",
		"*/",
		"func (t *testInterface) Buzz() {",
		"	panic(\"Not implemented\")",
		"}")

	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).Generate(&actual)
	assert.NoError(t, err)

	AssertCodeIsSame(t, expected, actual.String())
}
