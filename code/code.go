package code

import (
	"os"
	"text/template"
)

type Generator interface {
	ArgsError() error

	Cmd() (string, string)

	AssertArgs(args []string) bool

	Template() ([]byte, error)

	Params(common *CommonParams, args []string) interface{}
}

func Generate(cmd Generator, params *CommonParams, types []string) error {
	if cmd.AssertArgs(types) {
		return cmd.ArgsError()
	}

	p := cmd.Params(params, types)

	body, err := cmd.Template()
	if err != nil {
		return err
	}

	tmpl, err := template.New("t").Parse(string(body))
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, p)
}
