// itertools is an experimental another view in iterators. Gives more functions but does not allow to return errors
package itertools

type T0 string

type T0BinaryOp func(x T0, y T0) T0

type T0Predicate func(i int, item T0) bool

type T0Channel <-chan T0

func FromT0Slice(slice []T0) T0Channel {
	iter := make(chan T0)

	go func() {
		for _, it := range slice {
			iter <- it
		}
		close(iter)
	}()

	return iter
}

// Collect collects items from the iterator to a slice and returns it
func (c T0Channel) Collect() []T0 {
	res := []T0{}

	for it := range c {
		res = append(res, it)
	}
	return res
}

// Filter returns a channel and fill it with filtered results
func (c T0Channel) Filter(f T0Predicate) T0Channel {
	out := make(chan T0)

	go func() {
		i := -1
		for v := range c {
			i++
			if !f(i, v) {
				continue
			}
			out <- v
		}
		close(out)
	}()

	return T0Channel(out)
}

// Count return the number of items in the channel
func (c T0Channel) Count() int {
	i := 0
	for _ = range c {
		i++
	}
	return i
}

// Any returns true if predicate is true at leas for one item in the channel
func (c T0Channel) Any(f T0Predicate) bool {
	i := -1
	for v := range c {
		i++
		if f(i, v) {
			return true
		}
	}
	return false
}

// All returns true if predicate is true for all of the items in the channel
func (c T0Channel) All(f T0Predicate) bool {
	i := -1
	for v := range c {
		i++
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Reduce is a typical reduce operation
func (c T0Channel) Reduce(op T0BinaryOp) T0 {
	var res T0
	first := true

	for v := range c {
		if first {
			res = v
			first = false
			continue
		}
		res = op(res, v)
	}
	return res
}
