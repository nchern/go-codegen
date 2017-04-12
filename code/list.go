package code

import "fmt"

type ListGenerator struct{}

func (c *ListGenerator) ArgsError() error {
	return fmt.Errorf("Please provide <value-type>")
}

func (c *ListGenerator) AssertArgs(args []string) bool {
	return len(args) != 1
}

func (c *ListGenerator) Template() ([]byte, error) {
	return Asset("templates/list.go.t")
}

func (c *ListGenerator) Cmd() (cmd string, shortDescription string) {
	cmd = "list"
	shortDescription = "Generates typed list: []<value-type>"
	return
}

func (c *ListGenerator) Params(common *CommonParams, types []string) interface{} {
	common.AddTypeName(types[0])
	return common
}
