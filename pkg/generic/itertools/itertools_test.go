package itertools

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterAndCollectShouldWork(t *testing.T) {
	var tests = []struct {
		name      string
		predicate T0Predicate
		given     []T0
		expected  []T0
	}{
		{"1",
			func(i int, s T0) bool { return !strings.HasPrefix(string(s), "b") },
			[]T0{"foo", "bar", "buzz"},
			[]T0{"foo"}},
		{"2",
			func(i int, s T0) bool { return i > 0 },
			[]T0{"foo", "bar", "buzz"},
			[]T0{"bar", "buzz"}},

		{"3",
			func(i int, s T0) bool { return false },
			[]T0{"foo", "bar", "buzz"},
			[]T0{}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := FromT0Slice(tt.given).Filter(tt.predicate).Collect()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestAnyShould(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	var tests = []struct {
		name      string
		expected  bool
		predicate T0Predicate
	}{
		{"return element from the middle",
			true, func(i int, x T0) bool { return x == "bar" }},
		{"return nothing on false predicate",
			false, func(i int, x T0) bool { return x == "foobar" }},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := FromT0Slice(src).Any(tt.predicate)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestBothAnyAndAllShouldFullyDrainChannel(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}
	ch := FromT0Slice(src)
	actual := ch.Any(func(_ int, x T0) bool { return x == "foo" })
	assert.True(t, actual)
	for range ch {
		panic("must never be called")
	}

	ch = FromT0Slice(src)
	actual = ch.All(func(_ int, x T0) bool { return false })
	assert.False(t, actual)
	for range ch {
		panic("must never be called")
	}
}

func TestAllShould(t *testing.T) {
	src := []T0{"foo", "bar", "buzz"}

	var tests = []struct {
		name      string
		expected  bool
		predicate T0Predicate
	}{
		{"be true on true",
			true, func(i int, x T0) bool { return len(x) < 5 }},
		{"be false if predicate returns false at least onece",
			false, func(i int, x T0) bool { return strings.HasPrefix(string(x), "b") }},
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
