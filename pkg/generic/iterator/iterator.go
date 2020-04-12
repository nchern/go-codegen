// Package iterator provides a built-in implementation of an ideomatic generic iterator
package iterator

// T0 is a generic type variable placeholder of an iterator element. It will not appear in the generated code
type T0 string

// T0GeneratorFunc is a function that should generate elements and send them to a given channel
type T0GeneratorFunc func(generator chan<- T0) error

// T0SliceGenerator generates elements from a given slice
func T0SliceGenerator(src []T0) T0GeneratorFunc {
	return func(iter chan<- T0) error {
		for _, v := range src {
			iter <- v
		}
		return nil
	}
}

// T0Iter implements iterator over T0 type elements
type T0Iter interface {
	// Err returns error if it happened during generation
	Err() error

	// Next returns next element from iterator
	Next() <-chan T0
}

type iter struct {
	err  error
	iter chan T0
}

// Err returns error if it happened during generation
func (i *iter) Err() error {
	return i.err
}

// Next returns next element from iterator
func (i *iter) Next() <-chan T0 {
	return i.iter
}

// Generate creates Iterator from generator func
func Generate(f T0GeneratorFunc) T0Iter {
	gen := &iter{
		iter: make(chan T0),
	}

	go func() {
		gen.err = f(gen.iter)
		close(gen.iter)
	}()

	return gen
}
