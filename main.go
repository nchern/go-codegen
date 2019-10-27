package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nchern/go-codegen/pkg/code"
	"github.com/nchern/go-codegen/pkg/constructor"
	"github.com/nchern/go-codegen/pkg/generic"
	"github.com/nchern/go-codegen/pkg/immutable"
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

var (
	pkgName  = ""
	filename = ""

	buildInType = ""

	commands = []*cobra.Command{
		{
			Use:   "generic [list of concrete types to substitute type vars]",
			Short: "Processes go source as generic file and outputs code with substituted type vars",
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				var err error
				var generator generic.Generator
				if buildInType != "" {
					generator, err = generic.BuiltIn(buildInType)
					dieIf(err)
				} else {
					generator = generic.FromFile(filename)
				}
				err = code.WrapWithBannerPrinter(
					generator.
						WithPackageName(pkgName).
						WithTypeMapping(generic.TypeMapFromStrings(args...))).
					Generate(os.Stdout)
				dieIf(err)
			},
		},
		{
			Use:   "immutable",
			Short: "Generates immutable implementation by a given interface.",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err := code.WrapWithBannerPrinter(
					immutable.FromFile(filename).
						WithPackageName(pkgName)).
					Generate(os.Stdout)
				dieIf(err)
			},
		},
		{
			Use:   "struct",
			Short: "Generates initializer func for a given struct read from stdin",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err := constructor.FromReader(os.Stdin).
					WithPackageName(pkgName).
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

	rootCmd.PersistentFlags().StringVarP(&pkgName, "pkg", "p", "", "Golang package name. Substitutes existing package name or makes generator to add one")
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "", "input file name (reqiured, if no built-ins used)")

	commands[0].Flags().StringVarP(&buildInType, "type", "t", "",
		fmt.Sprintf("Generates based on predefined generic file. One of: %s", strings.Join(generic.BuiltInTypes(), ", ")))

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	err := rootCmd.Execute()
	dieIf(err)
}
