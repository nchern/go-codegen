package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nchern/go-codegen/pkg/code"
	"github.com/nchern/go-codegen/pkg/constructor"
	"github.com/nchern/go-codegen/pkg/generic"
	"github.com/nchern/go-codegen/pkg/immutable"
	"github.com/nchern/go-codegen/pkg/impl"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)
}

func dieIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func inputStream(shouldMirrorSource bool) io.Reader {
	if shouldMirrorSource {
		return io.TeeReader(os.Stdin, os.Stdout)
	}
	return os.Stdin
}

var (
	flagPkgName  = ""
	flagFileName = ""

	flagBuiltInType = ""

	flagOutputSource = false

	commands = []*cobra.Command{
		{
			Use:   "generic [list of concrete types to substitute type vars]",
			Short: "Processes go source as generic file and outputs code with substituted type vars.",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				var err error
				var generator generic.Generator
				if flagBuiltInType != "" {
					generator, err = generic.BuiltIn(flagBuiltInType)
					dieIf(err)
				} else {
					generator = generic.FromFile(flagFileName)
				}
				err = code.WrapWithBannerPrinter(
					generator.
						WithPackageName(flagPkgName).
						WithTypeMapping(generic.TypeMapFromStrings(args...))).
					Generate(os.Stdout)
				dieIf(err)
			},
		},
		{
			Use:   "immutable",
			Short: "Generates immutable structure implementation by a given interface.",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err := code.WrapWithBannerPrinter(
					immutable.FromFile(flagFileName).
						WithPackageName(flagPkgName)).
					Generate(os.Stdout)
				dieIf(err)
			},
		},
		{
			Use:   "constructor",
			Short: "Generates constructor function for a given struct read from stdin.",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err := constructor.FromReader(inputStream(flagOutputSource)).
					Generate(os.Stdout)
				dieIf(err)
			},
		},
		{
			Use:   "impl",
			Short: "Generates minimal interface implementation.",
			Long: "Generates stubs to implement a given interface\n" +
				"The difference between this command and impl(https://github.com/josharian/impl) utility is that interface declaration is read from stdin.\n" +
				"So that it's really easy to use it with editors like Vim",
			Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err := impl.FromReader(inputStream(flagOutputSource)).
					Generate(os.Stdout)
				dieIf(err)
			},
		},
	}
)

func main() {
	rootCmd := &cobra.Command{
		Use:  "go-codegen",
		Long: "Go code generation tool. Prints output to stdout",
	}

	rootCmd.PersistentFlags().StringVarP(&flagPkgName, "pkg", "p", "", "Golang package name. Substitutes existing package name or makes generator to add one")
	rootCmd.PersistentFlags().StringVarP(&flagFileName, "file", "f", "", "input file name (reqiured, if no built-ins used)")

	commands[0].Flags().StringVarP(&flagBuiltInType, "type", "t", "",
		fmt.Sprintf("Generates based on predefined generic file. One of: %s", strings.Join(generic.BuiltInTypes(), ", ")))

	commands[2].Flags().BoolVarP(&flagOutputSource, "out-src", "s", false, "if set, outputs the input source code before generated code. Could be helpful with editor integrations(e.g. vim)")
	commands[3].Flags().BoolVarP(&flagOutputSource, "out-src", "s", false, "if set, outputs the input source code before generated code. Could be helpful with editor integrations(e.g. vim)")

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	err := rootCmd.Execute()
	dieIf(err)
}
