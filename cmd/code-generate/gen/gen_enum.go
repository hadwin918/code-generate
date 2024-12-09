package gen

import (
	"github.com/go-courier/enumeration/generator"
	"github.com/go-courier/packagesx"
	"github.com/hadwin918/code-generate/internal/generate"
	"github.com/spf13/cobra"
)

func init() {
	CmdGen.AddCommand(cmdGenEnum)
}

var cmdGenEnum = &cobra.Command{
	Use:   "enum",
	Short: "generate interfaces of enumeration",
	Run: func(cmd *cobra.Command, args []string) {
		generate.RunGenerator(func(pkg *packagesx.Package) generate.Generator {
			g := generator.NewGenerator(pkg)
			g.Scan(args...)
			return g
		})
	},
}
