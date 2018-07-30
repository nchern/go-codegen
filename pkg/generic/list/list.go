package list

import (
	"sort"
	"sync"
)

type T0 string

type T0Predicate func(T0) bool

type T0Visitor func(i int, val T0) bool

type T0LessFunc func(first, second T0) bool

type T0List struct {
	list   []T0
	mutex  sync.RWMutex
	isSync bool
}

func NewT0List() *T0List {
	return NewT0ListFromSlice([]T0{}...)
}

func NewT0ListFromSlice(items ...T0) *T0List {
	l := make([]T0, len(items))
	copy(l, items)
	return &T0List{list: l}
}

func NewSyncronizedT0List(items ...T0) *T0List {
	res := NewT0ListFromSlice(items...)
	res.isSync = true
	return res
}

func (l *T0List) Filter(f T0Predicate) *T0List {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	result := &T0List{list: []T0{}, isSync: l.isSync}
	for _, v := range l.list {
		if f(v) {
			result.list = append(result.list, v)
		}
	}
	return result
}

func (l *T0List) IFilter(f T0Predicate) <-chan T0 {
	results := make(chan T0)
	go func() {
		if l.isSync {
			l.mutex.RLock()
			defer l.mutex.RUnlock()
		}
		for _, v := range l.list {
			if f(v) {
				results <- v
			}
		}
		close(results)
	}()

	return results
}

func (l *T0List) Iter() <-chan T0 {
	results := make(chan T0)
	go func() {
		if l.isSync {
			l.mutex.RLock()
			defer l.mutex.RUnlock()
		}
		for _, v := range l.list {
			results <- v
		}
		close(results)
	}()

	return results
}

func (l *T0List) Each(visitor T0Visitor) *T0List {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	for i, v := range l.list {
		if !visitor(i, v) {
			return l
		}
	}
	return l
}

func (l *T0List) Get(i int) T0 {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	return l.list[i]
}

func (l *T0List) Any(f T0Predicate) bool {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	for _, v := range l.list {
		if f(v) {
			return true
		}
	}
	return false
}

func (l *T0List) All(f T0Predicate) bool {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	for _, v := range l.list {
		if !f(v) {
			return false
		}
	}
	return true
}

func (l *T0List) FindFirst(f T0Predicate, defaultVal T0) T0 {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	for _, v := range l.list {
		if f(v) {
			return v
		}
	}
	return defaultVal
}

func (l *T0List) FindLast(f T0Predicate, defaultVal T0) T0 {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	found := defaultVal
	for _, v := range l.list {
		if f(v) {
			found = v
		}
	}
	return found
}

// For sort

func (l *T0List) Len() int {
	return len(l.list)
}

func (l *T0List) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

func (l *T0List) Sort(byFunc T0LessFunc) *T0List {
	if l.isSync {
		l.mutex.Lock()
		defer l.mutex.Unlock()
	}

	sorter := &T0ListSorter{l, byFunc}
	sort.Sort(sorter)

	return l
}

func (l *T0List) Clone() *T0List {
	if l.isSync {
		l.mutex.RLock()
		defer l.mutex.RUnlock()
	}

	copied := make([]T0, len(l.list))
	copy(copied, l.list)

	r := NewT0ListFromSlice(copied...)
	r.isSync = l.isSync

	return r
}

// Mutators

func (l *T0List) Set(i int, val T0) {
	if l.isSync {
		l.mutex.Lock()
		defer l.mutex.Unlock()
	}

	l.list[i] = val
}

func (l *T0List) Append(items ...T0) {
	if l.isSync {
		l.mutex.Lock()
		defer l.mutex.Unlock()
	}

	l.list = append(l.list, items...)
}

func (l *T0List) Prepend(items ...T0) {
	if l.isSync {
		l.mutex.Lock()
		defer l.mutex.Unlock()
	}

	l.list = append(items, l.list...)
}

func (l *T0List) Pop(defaultVal T0) T0 {
	if l.isSync {
		l.mutex.Lock()
		defer l.mutex.Unlock()
	}

	if len(l.list) < 1 {
		return defaultVal
	}

	result := l.list[len(l.list)-1]

	l.list = l.list[:len(l.list)-1]

	return result
}

// PopRight

type T0ListSorter struct {
	*T0List
	lessFn T0LessFunc
}

func (s *T0ListSorter) Less(i, j int) bool {
	return s.lessFn(s.list[i], s.list[j])
}
