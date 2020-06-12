package itertools

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldFilterAndCollect(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	ch := FromT0Slice(src)

	filtered := ch.Filter(func(i int, s T0) bool { return !strings.HasPrefix(string(s), "b") })
	for v := range filtered {
		assert.Equal(t, T0("foo"), v)
	}

	filtered = FromT0Slice(src).Filter(func(i int, s T0) bool { return i > 0 })
	actual := filtered.Collect()
	assert.Equal(t, []T0{"bar", "buzz"}, actual)

	filtered = FromT0Slice(src).Filter(func(i int, s T0) bool { return false })
	actual = filtered.Collect()
	assert.Len(t, actual, 0)
}

func TestShouldAnyBeOK(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	var tests = []struct {
		name      string
		expected  bool
		predicate T0Predicate
	}{
		{"true", true, func(i int, x T0) bool { return x == "bar" }},
		{"false", false, func(i int, x T0) bool { return x == "foobar" }},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := FromT0Slice(src).Any(tt.predicate)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestShouldAllBeOK(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	var tests = []struct {
		name      string
		expected  bool
		predicate T0Predicate
	}{
		{"true", true, func(i int, x T0) bool { return len(x) < 5 }},
		{"false", false, func(i int, x T0) bool { return strings.HasPrefix(string(x), "b") }},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := FromT0Slice(src).All(tt.predicate)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestShouldReduce(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	var tests = []struct {
		name     string
		expected T0
		given    []T0
	}{
		{"concat", T0("foobarbuzz"), src},
		{"single element", T0("foo"), []T0{"foo"}},
		{"empty slice", T0(""), []T0{}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := FromT0Slice(tt.given).Reduce(func(a, b T0) T0 {
				return a + b
			})
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestShouldCount(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	actual := FromT0Slice(src).Count()
	assert.Equal(t, 3, actual)
}
