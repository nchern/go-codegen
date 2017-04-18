{{.Pkg}}

type {{.Name}}Predicate func ({{.T0}}) bool

type {{.Name}}Visitor func (i int, val {{.T0}}) bool

type {{.Name}}LessFunc func (first, second {{.T0}}) bool

type {{.Name}} struct {
    _list []{{.T0}}
    lessFn {{.Name}}LessFunc

    {{if .IsSync}}
    _lock sync.RWMutex
    {{end}}
}

func New{{.Name}}() *{{.Name}} {
    return New{{.Name}}FromSlice([]{{.T0}}{})
}

func New{{.Name}}FromSlice(items []{{.T0}}) *{{.Name}} {
    return &{{.Name}}{ _list: items }
}

func (l *{{.Name}}) Filter(f {{.Name}}Predicate) *{{.Name}} {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    result := &{{.Name}}{ _list: []{{.T0}}{} }
    for _, v := range l._list {
        if f(v) {
            result._list = append(result._list, v)
        }
    }
    return result
}

func (l *{{.Name}}) IFilter(f {{.Name}}Predicate) <-chan {{.T0}}  {
    results := make(chan {{.T0}})
    go func() {
        {{if .IsSync}}
        l._lock.RLock()
        defer l._lock.RUnlock()
        {{end}}
        for _, v := range l._list {
            if f(v) {
                results <- v
            }
        }
        close(results)
    }()

    return results
}

func (l *{{.Name}}) Iter() <-chan {{.T0}}  {
    results := make(chan {{.T0}})
    go func() {
        {{if .IsSync}}
        l._lock.RLock()
        defer l._lock.RUnlock()
        {{end}}
        for _, v := range l._list {
            results <- v
        }
        close(results)
    }()

    return results
}

func (l *{{.Name}}) Each(visitor {{.Name}}Visitor) *{{.Name}} {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    for i, v := range l._list {
        if !visitor(i, v) {
            return l
        }
    }
    return l
}


func (l *{{.Name}}) Any(f {{.Name}}Predicate) bool {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    for _, v := range l._list {
        if f(v) {
            return true
        }
    }
    return false
}

func (l *{{.Name}}) All(f {{.Name}}Predicate) bool {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    for _, v := range l._list {
        if !f(v) {
            return false
        }
    }
    return true
}

func (l *{{.Name}}) FindFirst(f {{.Name}}Predicate, defaultVal {{.T0}}) {{.T0}} {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    for _, v := range l._list {
        if f(v) {
            return v
        }
    }
    return defaultVal
}

func (l *{{.Name}}) FindLast(f {{.Name}}Predicate, defaultVal {{.T0}}) {{.T0}} {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    found := defaultVal
    for _, v := range l._list {
        if f(v) {
            found = v
        }
    }
    return found
}


// For sort

func (l *{{.Name}}) ByFunc(lessFn {{.Name}}LessFunc) *{{.Name}}{
    {{if .IsSync}}
    l._lock.Lock()
    defer l._lock.Unlock()
    {{end}}

    l.lessFn = lessFn
    return l
}

func (l *{{.Name}}) Less(i, j int) bool{
    return l.lessFn(l._list[i], l._list[j])
}


func (l *{{.Name}}) Len() int {
    return len(l._list)
}

func (l *{{.Name}}) Swap(i, j int) {
    l._list[i], l._list[j] = l._list[j], l._list[i]
}

func (l *{{.Name}}) Sort() {
    {{if .IsSync}}
    l._lock.Lock()
    defer l._lock.Unlock()
    {{end}}

    sort.Sort(l)

}


func (l *{{.Name}}) Clone() *{{.Name}} {
    {{if .IsSync}}
    l._lock.RLock()
    defer l._lock.RUnlock()
    {{end}}

    copied := make([]{{.T0}}, len(l._list))
    copy(copied, l._list)

    return New{{.Name}}FromSlice(copied)
}

// Mutators

func (l *{{.Name}}) Append(items ...{{.T0}}) {
    {{if .IsSync}}
    l._lock.Lock()
    defer l._lock.Unlock()
    {{end}}

    l._list = append(l._list, items...)
}

func (l *{{.Name}}) Prepend(items ...{{.T0}}) {
    {{if .IsSync}}
    l._lock.Lock()
    defer l._lock.Unlock()
    {{end}}

    l._list = append(items, l._list...)
}

func (l *{{.Name}}) Pop(defaultVal {{.T0}}) {{.T0}} {
    {{if .IsSync}}
    l._lock.Lock()
    defer l._lock.Unlock()
    {{end}}

    if len(l._list) < 1 {
        return defaultVal
    }

    result := l._list[len(l._list)-1]

    l._list = l._list[:len(l._list)-1]

    return result
}
// PopRight 
