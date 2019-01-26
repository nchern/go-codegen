package model

// DTO is the source to generate immutable value object.
// Must be an interface with methods that accept no params and return a single value only.
type DTO interface {
	Foo() string
	Bar() int
	Buzz() bool

	Value() float64
}
