package code

type CommonParams struct {
	IsSync  bool
	Name    string
	PkgName string

	T []string
}

func (p *CommonParams) Pkg() string {
	if p.PkgName != "" {
		return "package " + p.PkgName
	}
	return ""
}

func (p *CommonParams) AddTypeName(name ...string) {
	p.T = append(p.T, name...)
}

func (p *CommonParams) T0() string {
	return p.T[0]
}

func (p *CommonParams) T1() string {
	return p.T[1]
}

func (p *CommonParams) T2() string {
	return p.T[2]
}
