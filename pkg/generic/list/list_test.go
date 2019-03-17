package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedArr = []T0{"buzz", "foo", "bar"}
)

func TestListClone(t *testing.T) {

	src := NewT0ListFromSlice(expectedArr...)
	inTest := src.Clone()

	for i, expectedVal := range expectedArr {
		assert.Equal(t, expectedVal, inTest.Get(i))
	}

	v := T0("FOO")
	inTest.Set(1, v)
	assert.Equal(t, v, inTest.Get(1))

	assert.Equal(t, expectedArr[1], src.Get(1))
}

func TestListBasics(t *testing.T) {
	inTest := NewT0List()
	assert.Equal(t, 0, inTest.Len())

	inTest.Append("x")
	assert.Equal(t, 1, inTest.Len())
	assert.Equal(t, T0("x"), inTest.Get(0))

	inTest.Prepend("y")
	assert.Equal(t, 2, inTest.Len())
	assert.Equal(t, T0("y"), inTest.Get(0))

	defaultVal := T0("")
	assert.Equal(t, T0("x"), inTest.Pop(defaultVal))
	assert.Equal(t, T0("y"), inTest.Pop(defaultVal))
	assert.Equal(t, 0, inTest.Len())
	assert.Equal(t, defaultVal, inTest.Pop(defaultVal))
}

func TestListIterations(t *testing.T) {
	inTest := NewT0ListFromSlice(expectedArr...)
	i := 0
	for v := range inTest.Iter() {
		assert.Equal(t, expectedArr[i], v)
		i++
	}
	for v := range inTest.IFilter(func(v T0) bool { return v == expectedArr[2] }) {
		assert.Equal(t, expectedArr[2], v)
	}

	assert.True(t, inTest.Any(func(v T0) bool { return v == expectedArr[1] }))
	assert.False(t, inTest.Any(func(v T0) bool { return v == T0("") }))

	assert.True(t, inTest.All(func(v T0) bool { return v != T0("ac") }))
	assert.False(t, inTest.All(func(v T0) bool { return v == expectedArr[1] }))
}

func TestSort(t *testing.T) {
	inTest := NewT0ListFromSlice(expectedArr...)
	inTest.Sort(func(a, b T0) bool { return a < b })

	expected := []T0{"bar", "buzz", "foo"}

	for i, expectedVal := range expected {
		assert.Equal(t, expectedVal, inTest.Get(i))
	}
}
