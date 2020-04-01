package iterator

type T0 string

type T0GeneratorFunc func(generator chan<- T0) error

func T0SliceGenerator(src []T0) T0GeneratorFunc {
	return func(iter chan<- T0) error {
		for _, v := range src {
			iter <- v
		}
		return nil
	}
}

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
