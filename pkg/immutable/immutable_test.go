package immutable

import (
	"bytes"
	"os"
	"strings"
	"testing"

	. "github.com/nchern/go-codegen/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

const (
	generatedImmutableSrc = `type valueStruct  struct {
		SomeNameField string 'json:"some_name"'
		SomeValueField int 'json:"some_value"'
	}
	func (m *valueStruct) SomeName() string { return m.SomeNameField }
	func (m *valueStruct) SomeValue() int { return m.SomeValueField }

	type ImmutableValueBuilder struct { value *valueStruct }
	// NewValueBuilder creates new ImmutableValueBuilder builder
	func NewValueBuilder() *ImmutableValueBuilder { return &ImmutableValueBuilder{ &valueStruct{} } }
	// SomeName sets corresponding field value
	func (b *ImmutableValueBuilder) SomeName(SomeNameField string) *ImmutableValueBuilder { b.value.SomeNameField = SomeNameField; return b }
	// SomeValue sets corresponding field value
	func (b *ImmutableValueBuilder) SomeValue(SomeValueField int) *ImmutableValueBuilder { b.value.SomeValueField = SomeValueField; return b }

	// Build builds the immutable structure
	func (b ImmutableValueBuilder) Build() Value { ret := *b.value; return &ret }`
)

func TestCamelCaseToSnakeCase(t *testing.T) {
	assert.Equal(t,
		"name",
		camelCaseToSnakeCase("Name"))
	assert.Equal(t,
		"n",
		camelCaseToSnakeCase("N"))
	assert.Equal(t,
		"foo_bar",
		camelCaseToSnakeCase("FooBar"))
	assert.Equal(t,
		"f_bar",
		camelCaseToSnakeCase("FBar"))
	assert.Equal(t,
		"foo_b",
		camelCaseToSnakeCase("fooB"))
	assert.Equal(t,
		"f_b",
		camelCaseToSnakeCase("FB"))
	assert.Equal(t,
		"foo_bar_buzz",
		camelCaseToSnakeCase("FooBarBuzz"))
}

func TestMethodStruct(t *testing.T) {
	inTest := &method{Name: "Foo", RetValue: "int64"}

	assert.Equal(t, "FooField", inTest.FieldName())
	assert.Equal(t, "FooField int64 `json:\"foo\"`", inTest.StructField())

	assert.Equal(t, "func (m *Obj) Foo() int64 { return m.FooField }",
		inTest.generateImmutableSetter("Obj"))

	assert.Equal(t, "// Foo sets corresponding field value\nfunc (b *ObjBuilder) Foo(FooField int64) *ObjBuilder { b.value.FooField = FooField; return b }",
		inTest.generateBuilderSetter("ObjBuilder"))
}

func TestTypeInfo(t *testing.T) {
	inTest := typeInfo{Name: "Bar"}

	assert.Equal(t, "barStruct", inTest.StructName())
	assert.Equal(t, "ImmutableBarBuilder", inTest.BuilderName())

}

func TestShouldFailIfImmutableMethodHasParams(t *testing.T) {
	srcText := `package immutable
	type Value interface {
		Foo(a int) string
	}`
	file := CreateGoFile(srcText)
	defer os.Remove(file.Name())

	buf := bytes.Buffer{}
	err := FromFile(file.Name()).Generate(&buf)
	assert.Error(t, err)
	assert.Equal(t, ErrUnsupportedMethodSignature, err)
}

func TestShouldFailIfImmutableMethodHasMoreThanOneReturnParams(t *testing.T) {
	srcText := `package immutable
	type Value interface {
		Foo() (string, int)
	}`
	file := CreateGoFile(srcText)
	defer os.Remove(file.Name())

	buf := bytes.Buffer{}
	err := FromFile(file.Name()).Generate(&buf)
	assert.Error(t, err)
	assert.Equal(t, ErrUnsupportedMethodSignature, err)
}

func TestShouldGenerateImmutableStructAndBuilder(t *testing.T) {
	srcText := `package immutable

	type Value interface {
		SomeName() string
		SomeValue() int
	}`

	file := CreateGoFile(srcText)
	defer os.Remove(file.Name())

	buf := bytes.Buffer{}
	err := FromFile(file.Name()).Generate(&buf)
	assert.NoError(t, err)

	expectedSrc := strings.Replace(generatedImmutableSrc, "'", "`", -1)
	actualSrc := strings.Trim(buf.String(), "\n")
	AssertCodeIsSame(t, expectedSrc, actualSrc)
}
