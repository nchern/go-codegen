{{.Pkg}}

type {{.Name}}Visitor func ({{.T0}}, {{.T1}}) bool

type {{ .Name }} struct {
    _map map[{{.T0}}]{{.T1}}

    {{if .IsSync}}
    _lock *sync.RWMutex
    {{end}}
}

func New{{.Name}}() *{{.Name}} {
    res := &{{.Name}}{
        _map: map[{{.T0}}]{{.T1}}{},
        {{if .IsSync}}
        _lock: &sync.RWMutex{},
        {{end}}
    }
    return res
}

func (m *{{.Name}}) Get(key {{.T0}}) (v {{.T1}}, found bool) {
    {{if .IsSync}} 
    m._lock.RLock()
    defer m._lock.RUnlock()
    {{end}}

    v, found = m._map[key]
    return
}

func (m *{{.Name}}) Set(key {{.T0}}, val {{.T1}}) {
    {{if .IsSync}} 
    m._lock.Lock()
    defer m._lock.Unlock()
    {{end}}

    m._map[key] = val
}

func (m *{{.Name}}) Each(visitor {{.Name}}Visitor) {
    {{if .IsSync}} 
    m._lock.RLock()
    defer m._lock.RUnlock()
    {{end}}

    for k, v := range m._map {
        if !visitor(k, v) { return }
    }
}

func (m *{{.Name}}) Update(src map[{{.T0}}]{{.T1}}) *{{.Name}}{
    {{if .IsSync}} 
    m._lock.Lock()
    defer m._lock.Unlock()
    {{end}}

    for k, v := range src {
        m._map[k] = v
    }
    return m
}

func (m *{{.Name}}) Remove(key {{.T0}}) {
    {{if .IsSync}} 
    m._lock.Lock()
    defer m._lock.Unlock()
    {{end}}

    delete(m._map, key)
}
