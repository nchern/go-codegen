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
	// T0 generic type name
	T0 TypeVar = "T0"
	// T1 generic type name
	T1 TypeVar = "T1"
	// T2 generic type name
	T2 TypeVar = "T2"
	// T3 generic type name
	T3 TypeVar = "T3"
	// T4 generic type name
	T4 TypeVar = "T4"
	// T5 generic type name
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

func (m TypeMap) substituteTypeVar(text textChunk) {
	for tVar, tVal := range m {
		if !strings.Contains(text.String(), string(tVar)) {
			continue
		}
		subs := tVal

		if toks := strings.Split(subs, "."); len(toks) > 1 {
			if len(toks) == 2 {
				// strip out package name from an exported identifier, like foo.Bar
				subs = toks[1]
			} else {
				panic(fmt.Errorf("bad identifier name: %s", tVal))
			}
		}

		if strings.Contains(tVal, "*") {
			subs = strings.Replace(subs, "*", "", -1) + "Ptr"
		}
		if strings.HasSuffix(tVal, "interface{}") {
			subs = strings.Replace(subs, "interface{}", "Object", -1)
		}
		if strings.HasPrefix(tVal, "[]") {
			subs = strings.TrimPrefix(subs, "[]") + "Slice"
		}

		text.SetString(strings.Replace(text.String(), string(tVar), strings.Title(subs), -1))
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

// Generator implements a generic code generator
type Generator struct {
	typeVars TypeMap

	src         interface{}
	filename    string
	packageName string
}

// FromFile returns Generator from a given file
func FromFile(filename string) *Generator {
	return &Generator{
		src:      nil,
		filename: filename,
	}
}

// FromBytes returns Generator from byte source
func FromBytes(src []byte) *Generator {
	return &Generator{
		src: src,
	}
}

// WithPackageName sets the package name
func (g *Generator) WithPackageName(name string) *Generator {
	g.packageName = name
	return g
}

// WithTypeMapping sets type vars mapping to concrete types
func (g *Generator) WithTypeMapping(typeVars TypeMap) *Generator {
	g.typeVars = typeVars
	return g
}

// Generate generates implementation based on generic code and writes it to the Writer
func (g *Generator) Generate(w io.Writer) error {
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
				// Ident is a generic type
				return true
			}
			g.typeVars.substituteTypeVar(&identText{n})
		case *ast.Comment:
			g.typeVars.substituteTypeVar(&commentText{n})
		}
		return true
	})

	if g.packageName != "" {
		node.Name = ast.NewIdent(g.packageName)
	}

	return printer.Fprint(w, fset, node)
}

type textChunk interface {
	String() string
	SetString(s string)
}

type identText struct {
	*ast.Ident
}

func (c *identText) String() string {
	return c.Name
}

func (c *identText) SetString(s string) {
	c.Name = s
}

type commentText struct {
	*ast.Comment
}

func (c *commentText) String() string {
	return c.Text
}

func (c *commentText) SetString(s string) {
	c.Text = s
}
