package impl

import (
	"bytes"
	"testing"

	"github.com/nchern/go-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateImplementation(t *testing.T) {
	source := `
type TestInterface interface {
	Foo() int
	Bar() CustomStruct
	FooBar() interface{}
	Fuzz() []*CustomStruct
}`

	expected := `
type testInterface struct {}

func (t *testInterface) Foo() int {
	panic("Not implemented")
}

func (t *testInterface) Bar() CustomStruct {
	panic("Not implemented")
}

func (t *testInterface) FooBar() interface{} {
	panic("Not implemented")
}

func (t *testInterface) Fuzz() []*CustomStruct {
	panic("Not implemented")
}
`
	var actual bytes.Buffer
	err := FromReader(bytes.NewBufferString(source)).Generate(&actual)
	assert.NoError(t, err)

	assert.Equal(t, testutil.FormatSrc(expected), testutil.FormatSrc(actual.String()))
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
