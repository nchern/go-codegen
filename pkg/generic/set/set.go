// Package set provides a built-in implementation of a generic set
package set

// T0 is a generic type variable placeholder of a set element. It will not appear in the generated code
type T0 string

// T0Set represents a set of T0 elements
type T0Set struct {
	m map[T0]bool
}

// NewT0Set creates an instance of T0Set filled with provided elements
func NewT0Set(vals ...T0) *T0Set {
	res := &T0Set{m: map[T0]bool{}}
	res.Add(vals...)
	return res
}

// Add adds an element to a set
func (s *T0Set) Add(vals ...T0) {
	for _, v := range vals {
		s.m[v] = true
	}
}

// Remove removes an element from a set
func (s *T0Set) Remove(x T0) {
	delete(s.m, x)
}

// Len returns elements number in the set
func (s *T0Set) Len() int {
	return len(s.m)
}

// Contains checks if a given element belongs to a set
func (s *T0Set) Contains(x T0) bool {
	return s.m[x]
}

// All returns a channel iterator over this set
func (s *T0Set) All() <-chan T0 {
	c := make(chan T0)
	go func() {
		for x := range s.m {
			c <- x
		}
		close(c)
	}()
	return c
}

// Union applies union operation to this set and a given set
func (s *T0Set) Union(b *T0Set) *T0Set {
	for x := range b.m {
		s.Add(x)
	}
	return s
}

// Subtract subtracts b from this set
func (s *T0Set) Subtract(b *T0Set) *T0Set {
	for x := range b.m {
		s.Remove(x)
	}
	return s
}

// Intersect finds intersection of y and this set
func (s *T0Set) Intersect(y *T0Set) *T0Set {
	a := s.m
	b := y.m
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make(map[T0]bool)
	for x := range a {
		if b[x] {
			res[x] = true
		}
	}
	s.m = res
	return s
}
