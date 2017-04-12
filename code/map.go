package code

import "fmt"

type MapGenerator struct{}

func (c *MapGenerator) ArgsError() error {
	return fmt.Errorf("Please provide <key-type> and <value-type>")
}

func (c *MapGenerator) AssertArgs(args []string) bool {
	return len(args) < 2
}

func (c *MapGenerator) Template() ([]byte, error) {
	return Asset("templates/map.go.t")
}

func (c *MapGenerator) Cmd() (cmd string, shortDescription string) {
	cmd = "map"
	shortDescription = "Generates typed map: map[<key-type>]<value-type>"
	return
}

func (c *MapGenerator) Params(common *CommonParams, types []string) interface{} {
	common.AddTypeName(types[0], types[1])
	return common
}
