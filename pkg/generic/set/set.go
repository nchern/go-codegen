package set

type T0 string

type T0Set struct {
	m map[T0]bool
}

func NewT0Set(vals ...T0) *T0Set {
	res := &T0Set{m: map[T0]bool{}}
	res.Add(vals...)
	return res
}

func (s *T0Set) Add(vals ...T0) {
	for _, v := range vals {
		s.m[v] = true
	}
}

func (s *T0Set) Remove(x T0) {
	delete(s.m, x)
}

func (s *T0Set) Len() int {
	return len(s.m)
}

func (s *T0Set) Contains(x T0) bool {
	return s.m[x]
}

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

func (s *T0Set) Union(b *T0Set) *T0Set {
	for x := range b.m {
		s.Add(x)
	}
	return s
}

func (s *T0Set) Subtract(b *T0Set) *T0Set {
	for x := range b.m {
		s.Remove(x)
	}
	return s
}

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
