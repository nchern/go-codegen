package impl

import (
	"bytes"
	"testing"

	"github.com/nchern/go-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateImplementation(t *testing.T) {
	source := text(
		"type TestInterface interface {",
		"	Foo(u *User) int",
		"	Bar(a int, b float64) CustomStruct",
		"	FooBar() interface{}",
		"	Fuzz() []*CustomStruct",
		"	IsGood() bool",
		"	FetchPtr() *Ptr",
		"}")

	expected := text(
		"type testInterface struct {}",
		"func (t *testInterface) Foo(u *User) int {",
		"	return 0",
		"}",
		"",
		"func (t *testInterface) Bar(a int, b float64) CustomStruct {",
		"	panic(\"Not implemented\")",
		"}",
		"",
		"func (t *testInterface) FooBar() interface{} {",
		"	return nil",
		"}",
		"",
		"func (t *testInterface) Fuzz() []*CustomStruct {",
		"	return nil",
		"}",
		"",
		"func (t *testInterface) IsGood() bool {",
		"	return false",
		"}",
		"",
		"func (t *testInterface) FetchPtr() *Ptr {",
		"	return nil",
		"}")

	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).Generate(&actual)
	assert.NoError(t, err)

	testutil.AssertCodeIsSame(t, expected, actual.String())
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
	source := text(
		"type TestInterface interface {",
		"	// Foo has a single line comment",
		"	Foo() int",

		"	// Bar has",
		"	// two single line comments",
		"	Bar(i int) string",

		"	/*",
		"		Buzz has multi-",
		"		line comment",
		"	*/",
		"	Buzz()",
		"}")

	expected := text(
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
		"	Buzz has multi-",
		"	line comment",
		"*/",
		"func (t *testInterface) Buzz() {",
		"	panic(\"Not implemented\")",
		"}")

	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).Generate(&actual)
	assert.NoError(t, err)

	testutil.AssertCodeIsSame(t, expected, actual.String())
}
