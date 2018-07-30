package iterator

type T0 string

type T0GeneratorFunc func(generator chan<- T0) error

type T0Iter interface {
	Err() error
	Next() <-chan T0
}

type iter struct {
	err  error
	iter chan T0
}

func (i *iter) Err() error {
	return i.err
}

func (i *iter) Next() <-chan T0 {
	return i.iter
}

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
