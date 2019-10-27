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

	expected = source + expected
	assert.Equal(t, strings.Trim(expected, "\n"), actual.String())
}
