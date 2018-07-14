package set

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicOperations(t *testing.T) {
	set := NewT0Set()
	assert.Equal(t, 0, set.Len())

	set.Add("foo", "bar", "buzz")
	assert.True(t, set.Contains("foo"))
	assert.True(t, set.Contains("bar"))
	assert.True(t, set.Contains("buzz"))
	assert.Equal(t, 3, set.Len())

	set.Remove("buzz")
	assert.False(t, set.Contains("buzz"))
}

func TestAllIteration(t *testing.T) {
	set := NewT0Set()

	set.Add("foo", "bar", "buzz")

	actual := []string{}
	for x := range set.All() {
		actual = append(actual, string(x))
	}
	sort.Strings(actual)
	expected := []string{"bar", "buzz", "foo"}
	assert.Equal(t, expected, actual)
}

func TestSetTheoryOperations(t *testing.T) {
	a := NewT0Set("foo", "bar", "buzz")
	b := NewT0Set("boo")

	a.Union(b)
	assert.Equal(t, 4, a.Len())
	assert.True(t, a.Contains("boo"))

	a.Subtract(NewT0Set("bar", "boo", "hello"))
	assert.Equal(t, 2, a.Len())
	assert.False(t, a.Contains("bar"))
	assert.False(t, a.Contains("boo"))

	a = NewT0Set("foo", "bar", "buzz", "xxx").Intersect(NewT0Set("foo", "xxx", "aaa", "bbb", "zzz"))
	assert.Equal(t, 2, a.Len())
	assert.True(t, a.Contains("foo"))
	assert.True(t, a.Contains("xxx"))
}
