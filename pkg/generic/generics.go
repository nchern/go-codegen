package generic

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"strings"
)

// TypeVar indicates generic type var
type TypeVar string

const (
	// T0
	T0 TypeVar = "T0"
	// T1
	T1 TypeVar = "T1"
	// T2
	T2 TypeVar = "T2"
	// T3
	T3 TypeVar = "T3"
	// T4
	T4 TypeVar = "T4"
	// T5
	T5 TypeVar = "T5"
)

var (
	allowedTypeVars = map[TypeVar]bool{
		T0: true,
		T1: true,
		T2: true,
		T3: true,
		T4: true,
		T5: true,
	}

	// ErrBadTypeVar indicates that type var is unsupported
	ErrBadTypeVar = errors.New("Bad type variable")
)

// TypeMap maps type vars
type TypeMap map[TypeVar]string

// TypeMapFromStrings returns TypeMap filled with given type var names
func TypeMapFromStrings(types ...string) TypeMap {
	res := TypeMap{}
	for i, v := range types {
		res[TypeVar(fmt.Sprintf("T%d", i))] = v
	}
	return res
}

func (m TypeMap) rewriteType(n *ast.Ident) bool {
	for tVar, tVal := range m {
		if TypeVar(n.Name) != tVar {
			continue
		}
		n.Name = tVal
		return true
	}
	return false
}

func (m TypeMap) substituteTypeVarInIdent(n *ast.Ident) {
	for tVar, tVal := range m {
		if !strings.Contains(n.Name, string(tVar)) {
			continue
		}
		subs := tVal
		if strings.Contains(tVal, "*") {
			subs = strings.Replace(subs, "*", "", -1) + "Ptr"
		}
		if strings.HasSuffix(tVal, "interface{}") {
			subs = strings.Replace(subs, "interface{}", "Object", -1)
		}
		if strings.HasPrefix(tVal, "[]") {
			subs = strings.TrimPrefix(subs, "[]") + "Slice"
		}

		n.Name = strings.Replace(n.Name, string(tVar), strings.Title(subs), -1)
	}
}

func (m TypeMap) stripTypeVarsDecls(node *ast.File) {
	indexesToRemove := map[int]bool{}
	for i, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}
			for _, s := range d.Specs {
				spec := s.(*ast.TypeSpec)
				for tVar := range m {
					if TypeVar(spec.Name.String()) == tVar {
						indexesToRemove[i] = true
					}
				}
			}
		}
	}
	decls := node.Decls
	node.Decls = nil
	for i, decl := range decls {
		if indexesToRemove[i] {
			continue
		}
		node.Decls = append(node.Decls, decl)
	}
}

// Generator represents abstract generic processor
type Generator interface {
	WithPackageName(name string) Generator
	WithTypeMapping(TypeMap) Generator
	Generate(io.Writer) error
}

type genericProcessor struct {
	typeVars TypeMap

	src         interface{}
	filename    string
	packageName string
}

// FromFile returns Generator from a given file
func FromFile(filename string) Generator {
	return &genericProcessor{
		src:      nil,
		filename: filename,
	}
}

// FromBytes returns Generator from byte source
func FromBytes(src []byte) Generator {
	return &genericProcessor{
		src: src,
	}
}

func (g *genericProcessor) WithPackageName(name string) Generator {
	g.packageName = name
	return g
}
func (g *genericProcessor) WithTypeMapping(typeVars TypeMap) Generator {
	g.typeVars = typeVars
	return g
}

func (g *genericProcessor) Generate(w io.Writer) error {
	for k := range g.typeVars {
		if !allowedTypeVars[k] {
			return ErrBadTypeVar
		}
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, g.filename, g.src, parser.ParseComments)
	if err != nil {
		return err
	}

	g.typeVars.stripTypeVarsDecls(node)

	ast.Inspect(node, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.Ident:
			if g.typeVars.rewriteType(n) {
				return true
			}
			g.typeVars.substituteTypeVarInIdent(n)
		}
		return true
	})

	if g.packageName != "" {
		node.Name = ast.NewIdent(g.packageName)
	}

	return printer.Fprint(w, fset, node)
}
