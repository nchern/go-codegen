// Package hashmap provides a built-in implementation of a generic map
package hashmap

import "sync"

// T0 is a generic type variable placeholder of a key type. It will not appear in the generated code
type T0 string

// T1 is a generic type variable placeholder of a value type. It will not appear in the generated code
type T1 int

// T0T1MapVisitor is a visitor function to visit map pairs
type T0T1MapVisitor func(T0, T1) bool

// T0T1Map exposes the contract of T0 to T1 map
type T0T1Map interface {
	// Each visits each element in the map. It stops iterations if visitor func returns false
	Each(visitor T0T1MapVisitor)

	// Get returns the value of a given key. If the key was not found the second return value will be false
	Get(key T0) (v T1, found bool)

	// Set sets the value of a given key
	Set(key T0, val T1)

	// Update updates the current map from a given map
	Update(src map[T0]T1) T0T1Map

	// Remove removes a given key from this map
	Remove(key T0) bool

	// Clone creates a copy of this map
	Clone() T0T1Map
}

type baseT0T1Map struct {
	_map map[T0]T1
}

// NewT0T1Map creates a basic instance of the T0T1Map. It is _unsafe_ for concurrent access.
func NewT0T1Map() T0T1Map {
	res := &baseT0T1Map{
		_map: map[T0]T1{},
	}
	return res
}

// NewT0T1MapSyncronized creates a concurrent safe instance of the T0T1Map
func NewT0T1MapSyncronized() T0T1Map {
	return &syncT0T1Map{
		inner: NewT0T1Map(),
	}
}

func (m *baseT0T1Map) Get(key T0) (v T1, found bool) {
	v, found = m._map[key]
	return
}

func (m *baseT0T1Map) Each(visitor T0T1MapVisitor) {
	for k, v := range m._map {
		if !visitor(k, v) {
			return
		}
	}
}

func (m *baseT0T1Map) Set(key T0, val T1) {
	m._map[key] = val
}

func (m *baseT0T1Map) Update(src map[T0]T1) T0T1Map {
	for k, v := range src {
		m._map[k] = v
	}
	return m
}

func (m *baseT0T1Map) Remove(key T0) bool {
	_, found := m._map[key]
	delete(m._map, key)

	return found
}

func (m *baseT0T1Map) Clone() T0T1Map {
	res := NewT0T1Map()
	for k, v := range m._map {
		res.Set(k, v)
	}

	return res
}

type syncT0T1Map struct {
	inner T0T1Map

	mutex sync.RWMutex
}

func (m *syncT0T1Map) Each(visitor T0T1MapVisitor) {
	m.mutex.RLock()
	m.inner.Each(visitor)
	m.mutex.RUnlock()
}

func (m *syncT0T1Map) Get(key T0) (v T1, found bool) {
	m.mutex.RLock()
	v, found = m.inner.Get(key)
	m.mutex.RUnlock()
	return
}

func (m *syncT0T1Map) Set(key T0, val T1) {
	m.mutex.Lock()
	m.inner.Set(key, val)
	m.mutex.Unlock()
}

func (m *syncT0T1Map) Update(src map[T0]T1) T0T1Map {
	m.mutex.Lock()
	m.inner.Update(src)
	m.mutex.Unlock()

	return m
}

func (m *syncT0T1Map) Remove(key T0) bool {
	m.mutex.Lock()
	found := m.inner.Remove(key)
	m.mutex.Unlock()

	return found
}

func (m *syncT0T1Map) Clone() T0T1Map {
	m.mutex.RLock()
	r := m.inner.Clone()
	m.mutex.RUnlock()
	return r
}
