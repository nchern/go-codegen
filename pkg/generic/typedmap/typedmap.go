package typedmap

import "sync"

type T0 string
type T1 int

type T0T1MapVisitor func(T0, T1) bool

type T0T1Map struct {
	_map map[T0]T1

	isSynced bool
	mutex    *sync.RWMutex
}

func NewT0T1Map() *T0T1Map {
	res := &T0T1Map{
		_map:  map[T0]T1{},
		mutex: &sync.RWMutex{},
	}
	return res
}

func NewT0T1MapSyncronized() *T0T1Map {
	res := NewT0T1Map()
	res.isSynced = true
	return res
}

func (m *T0T1Map) Get(key T0) (v T1, found bool) {
	if m.isSynced {
		m.mutex.RLock()
		defer m.mutex.RUnlock()
	}

	v, found = m._map[key]
	return
}

func (m *T0T1Map) Each(visitor T0T1MapVisitor) {
	if m.isSynced {
		m.mutex.RLock()
		defer m.mutex.RUnlock()
	}

	for k, v := range m._map {
		if !visitor(k, v) {
			return
		}
	}
}
func (m *T0T1Map) Set(key T0, val T1) {
	if m.isSynced {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}

	m._map[key] = val
}

func (m *T0T1Map) Update(src map[T0]T1) *T0T1Map {
	if m.isSynced {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}

	for k, v := range src {
		m._map[k] = v
	}
	return m
}

func (m *T0T1Map) Remove(key T0) bool {
	if m.isSynced {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}

	_, found := m._map[key]
	delete(m._map, key)

	return found
}

func (m *T0T1Map) Clone() *T0T1Map {
	if m.isSynced {
		m.mutex.Lock()
		defer m.mutex.Unlock()
	}

	res := NewT0T1Map()
	res.isSynced = m.isSynced
	for k, v := range m._map {
		res._map[k] = v
	}

	return res
}
