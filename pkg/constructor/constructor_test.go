package constructor

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateSimpleTypes(t *testing.T) {
	source := `type Foo struct {
	Bar string
	Bazz float64
//	FooBar []int
//	FooBarBazz []interface{}
//	Ptr *User
//	Friends []*User
}`

	expected := `
func NewFoo(bar string, bazz float64) *Foo {
	return &Foo{
		Bar: bar,
		Bazz: bazz,
	}
}`

	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).Generate(&actual)

	assert.NoError(t, err)
	assert.Equal(t, strings.Trim(expected, "\n"), actual.String())
}

func TestShouldGenerateWithOutputSource(t *testing.T) {
	source := `type Foo struct {}`

	expected := source + "\n" + `func NewFoo() *Foo {
	return &Foo{
	}
}`

	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).WithOutputSrc(true).Generate(&actual)

	assert.NoError(t, err)
	assert.Equal(t, strings.Trim(expected, "\n"), actual.String())
}

func TestShouldGenerateNothingOnUnsupportedTypes(t *testing.T) {
	var tests = []struct {
		name  string
		given string
	}{
		{"empty", "\n"},
		{"interface", "type Foo interface {}"},
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
