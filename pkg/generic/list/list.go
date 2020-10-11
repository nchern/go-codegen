// Package list provides a built-in implementation of a generic list
package list

import (
	"sort"
	"sync"
)

// T0 is a generic type variable placeholder of a key type. It will not appear in the generated code
type T0 string

// T0Predicate is a predicate function on T0 type
type T0Predicate func(T0) bool

// T0Visitor is a visitor function used to visit items of the list
type T0Visitor func(i int, val T0) bool

// T0LessFunc is a comparator function used to sort the list
type T0LessFunc func(first, second T0) bool

// T0List exposes a contract of a list of T0 elements
type T0List interface {
	// Filter returns a new list filled with items that are yielded true on given predicate
	Filter(f T0Predicate) T0List

	// IFilter is similar to Filter but returns a channel instead of a T0List instance
	IFilter(f T0Predicate) <-chan T0

	// Iter return a channel that is filled with this list elements
	Iter() <-chan T0

	// Each visits each element in the map. It stops iterations if visitor func returns false
	Each(visitor T0Visitor) T0List

	// Get returns i-th element from this list
	Get(i int) T0

	Any(f T0Predicate) bool
	All(f T0Predicate) bool
	FindFirst(f T0Predicate, defaultVal T0) T0
	FindLast(f T0Predicate, defaultVal T0) T0
	Len() int
	Swap(i, j int)
	Sort(byFunc T0LessFunc) T0List
	Clone() T0List
	Set(i int, val T0)
	Append(items ...T0)
	Prepend(items ...T0)
	Pop(defaultVal T0) T0
}

type baseT0List struct {
	list []T0
}

// NewT0List creates an empty list of T0 elements
func NewT0List() T0List {
	return NewT0ListFromSlice([]T0{}...)
}

// NewT0ListFromSlice creates a list of T0 element initialised from a given slice
func NewT0ListFromSlice(items ...T0) T0List {
	l := make([]T0, len(items))
	copy(l, items)
	return &baseT0List{list: l}
}

// NewSyncronizedT0List creates a concurrent safe instance of a T0List
func NewSyncronizedT0List(items ...T0) T0List {
	l := make([]T0, len(items))
	copy(l, items)
	inner := &baseT0List{list: l}
	return &syncT0List{inner: inner}
}

func (l *baseT0List) Filter(f T0Predicate) T0List {
	result := &baseT0List{list: []T0{}}
	for _, v := range l.list {
		if f(v) {
			result.list = append(result.list, v)
		}
	}
	return result
}

func (l *baseT0List) IFilter(f T0Predicate) <-chan T0 {
	results := make(chan T0)
	go func() {
		for _, v := range l.list {
			if f(v) {
				results <- v
			}
		}
		close(results)
	}()

	return results
}

func (l *baseT0List) Iter() <-chan T0 {
	results := make(chan T0)
	go func() {
		for _, v := range l.list {
			results <- v
		}
		close(results)
	}()

	return results
}

func (l *baseT0List) Each(visitor T0Visitor) T0List {
	for i, v := range l.list {
		if !visitor(i, v) {
			return l
		}
	}
	return l
}

func (l *baseT0List) Get(i int) T0 {
	return l.list[i]
}

func (l *baseT0List) Any(f T0Predicate) bool {
	for _, v := range l.list {
		if f(v) {
			return true
		}
	}
	return false
}

func (l *baseT0List) All(f T0Predicate) bool {
	for _, v := range l.list {
		if !f(v) {
			return false
		}
	}
	return true
}

func (l *baseT0List) FindFirst(f T0Predicate, defaultVal T0) T0 {
	for _, v := range l.list {
		if f(v) {
			return v
		}
	}
	return defaultVal
}

func (l *baseT0List) FindLast(f T0Predicate, defaultVal T0) T0 {
	found := defaultVal
	for _, v := range l.list {
		if f(v) {
			found = v
		}
	}
	return found
}

// For sort

func (l *baseT0List) Len() int {
	return len(l.list)
}

func (l *baseT0List) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
}

func (l *baseT0List) Sort(byFunc T0LessFunc) T0List {
	sorter := &T0ListSorter{l, byFunc}
	sort.Sort(sorter)

	return l
}

func (l *baseT0List) Clone() T0List {
	copied := make([]T0, len(l.list))
	copy(copied, l.list)

	r := NewT0ListFromSlice(copied...)

	return r
}

// Mutators

func (l *baseT0List) Set(i int, val T0) {
	l.list[i] = val
}

func (l *baseT0List) Append(items ...T0) {
	l.list = append(l.list, items...)
}

func (l *baseT0List) Prepend(items ...T0) {
	l.list = append(items, l.list...)
}

func (l *baseT0List) Pop(defaultVal T0) T0 {
	if len(l.list) < 1 {
		return defaultVal
	}

	result := l.list[len(l.list)-1]

	l.list = l.list[:len(l.list)-1]

	return result
}

// PopRight

type T0ListSorter struct {
	*baseT0List
	lessFn T0LessFunc
}

func (s *T0ListSorter) Less(i, j int) bool {
	return s.lessFn(s.list[i], s.list[j])
}

type syncT0List struct {
	inner *baseT0List

	mutex sync.RWMutex
}

func (l *syncT0List) Filter(f T0Predicate) T0List {
	l.mutex.RLock()
	r := l.inner.Filter(f)
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) IFilter(f T0Predicate) <-chan T0 {
	results := make(chan T0)
	go func() {
		l.mutex.RLock()
		for _, v := range l.inner.list {
			if f(v) {
				results <- v
			}
		}
		l.mutex.RUnlock()
		close(results)
	}()

	return results
}

func (l *syncT0List) Iter() <-chan T0 {
	results := make(chan T0)
	go func() {
		l.mutex.RLock()
		for _, v := range l.inner.list {
			results <- v
		}
		l.mutex.RUnlock()
		close(results)
	}()

	return results
}

func (l *syncT0List) Each(visitor T0Visitor) T0List {
	l.mutex.RLock()
	l.inner.Each(visitor)
	l.mutex.RUnlock()
	return l
}

func (l *syncT0List) Get(i int) T0 {
	l.mutex.RLock()
	v := l.inner.Get(i)
	l.mutex.RUnlock()
	return v
}

func (l *syncT0List) Any(f T0Predicate) bool {
	l.mutex.RLock()
	r := l.inner.Any(f)
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) All(f T0Predicate) bool {
	l.mutex.RLock()
	r := l.inner.All(f)
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) FindFirst(f T0Predicate, defaultVal T0) T0 {
	l.mutex.RLock()
	r := l.inner.FindFirst(f, defaultVal)
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) FindLast(f T0Predicate, defaultVal T0) T0 {
	l.mutex.RLock()
	r := l.inner.FindLast(f, defaultVal)
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) Len() int {
	l.mutex.RLock()
	r := l.inner.Len()
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) Swap(i int, j int) {
	l.mutex.Lock()
	l.inner.Swap(i, j)
	l.mutex.Unlock()
}

func (l *syncT0List) Sort(byFunc T0LessFunc) T0List {
	l.mutex.Lock()
	l.inner.Sort(byFunc)
	l.mutex.Unlock()
	return l
}

func (l *syncT0List) Clone() T0List {
	l.mutex.RLock()
	r := l.inner.Clone()
	l.mutex.RUnlock()
	return r
}

func (l *syncT0List) Set(i int, val T0) {
	l.mutex.Lock()
	l.inner.Set(i, val)
	l.mutex.Unlock()
}

func (l *syncT0List) Append(items ...T0) {
	l.mutex.Lock()
	l.inner.Append(items...)
	l.mutex.Unlock()
}

func (l *syncT0List) Prepend(items ...T0) {
	l.mutex.Lock()
	l.inner.Prepend(items...)
	l.mutex.Unlock()
}

func (l *syncT0List) Pop(defaultVal T0) T0 {
	l.mutex.Lock()
	r := l.inner.Pop(defaultVal)
	l.mutex.Unlock()
	return r
}
