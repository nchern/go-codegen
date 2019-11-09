package constructor

import (
	"bytes"
	"testing"

	"github.com/nchern/go-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateWithSimpleTypedFields(t *testing.T) {
	source := `
type Foo struct {
	Bar string
	Bazz float64
	Fuzz Phone
}`

	expected := `
func NewFoo(bar string, bazz float64, fuzz Phone) *Foo {
	return &Foo{
		Bar: bar,
		Bazz: bazz,
		Fuzz: fuzz,
	}
}
`

	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).Generate(&actual)

	assert.NoError(t, err)
	testutil.AssertCodeIsSame(t, expected, actual.String())
}

func TestShouldGenerateWitgComplextTypedFields(t *testing.T) {
	source := `
type Foo struct {
	Bar string
	Bazz interface{}
	FooBar []int
	FooBarBazz []interface{}
	Ptr *User
	Friends []*User
	Mapping map[string]interface{}
	PtrMapping map[float64]*User
}`

	expected := `
func NewFoo(bar string, bazz interface{}, fooBar []int, fooBarBazz []interface{}, ptr *User, friends []*User, mapping map[string]interface{}, ptrMapping map[float64]*User) *Foo {
	return &Foo{
		Bar: bar,
		Bazz: bazz,
		FooBar: fooBar,
		FooBarBazz: fooBarBazz,
		Ptr: ptr,
		Friends: friends,
		Mapping: mapping,
		PtrMapping: ptrMapping,
	}
}
`

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
