package typedmap

import "sync"

type T0 string
type T1 int

type T0T1MapVisitor func(T0, T1) bool

type T0T1Map interface {
	Each(visitor T0T1MapVisitor)
	Get(key T0) (v T1, found bool)
	Set(key T0, val T1)
	Update(src map[T0]T1) T0T1Map
	Remove(key T0) bool
	Clone() T0T1Map
}

type baseT0T1Map struct {
	_map map[T0]T1
}

func NewT0T1Map() T0T1Map {
	res := &baseT0T1Map{
		_map: map[T0]T1{},
	}
	return res
}

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
