package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nchern/go-codegen/pkg/generic"
	"github.com/nchern/go-codegen/pkg/immutable"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)
}

func failOnError(err error) {
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
				var processor generic.Processor
				if buildInType != "" {
					processor, err = generic.BuiltIn(buildInType, pkgName)
					failOnError(err)
				} else {
					processor = generic.FromFile(filename, pkgName)
				}
				err = processor.
					WithTypeMapping(generic.TypeMapFromStrings(args...)).
					WriteTo(os.Stdout)
				failOnError(err)
			},
		},
		{
			Use:   "immutable",
			Short: "Generates immutable implementation by a given interface.",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err := immutable.FromFile(filename).
					WriteTo(os.Stdout)
				failOnError(err)
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
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "", "input file name (reqiured)")

	commands[0].Flags().StringVarP(&buildInType, "type", "t", "",
		fmt.Sprintf("Generates based on predefined generic file. One of: %s", strings.Join(generic.BuiltInTypes(), ", ")))

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	failOnError(rootCmd.Execute())
}
