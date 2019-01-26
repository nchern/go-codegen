package main

import (
	"encoding/json"
	"testing"

	"github.com/nchern/go-codegen/tests/immutable/model"
	"github.com/stretchr/testify/assert"
)

func TestGeneratedImplementation(t *testing.T) {
	m := model.NewDTOBuilder().
		Foo("foo").
		Bar(101).
		Buzz(true).
		Value(3.14).
		Build()

	assert.Equal(t, "foo", m.Foo())
	assert.Equal(t, 101, m.Bar())
	assert.True(t, m.Buzz())
	assert.Equal(t, 3.14, m.Value())
}

func TestJSONEncodeDecode(t *testing.T) {
	expected := model.NewDTOBuilder().
		Foo("foo").
		Bar(101).
		Buzz(true).
		Value(3.14).
		Build()

	encoded, err := json.Marshal(expected)

	assert.NoError(t, err)
	assert.True(t, len(encoded) > 0)
	//	assert.Equal(t, "{d}", string(encoded))

	actual := model.NewDTOBuilder().Build()
	assert.NoError(t, json.Unmarshal(encoded, actual))

	assert.Equal(t, expected.Foo(), actual.Foo())
	assert.Equal(t, expected.Bar(), actual.Bar())
	assert.Equal(t, expected.Buzz(), actual.Buzz())
	assert.Equal(t, expected.Value(), actual.Value())
}
