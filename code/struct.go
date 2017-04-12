package code

import "errors"

type StructConv struct{}

func (g *StructConv) ArgsError() error {
	return errors.New("Please provide <target-struct> and <source-struct>")
}

func (g *StructConv) Cmd() (string, string) {
	panic("not implemented")
}

func (g *StructConv) AssertArgs(args []string) bool {
	return len(args) == 2
}

func (g *StructConv) Template() ([]byte, error) {
	panic("not implemented")
}

func (g *StructConv) Params(common *CommonParams, args []string) interface{} {
	panic("not implemented")
}
