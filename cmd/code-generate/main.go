package main

import (
	"fmt"
	"github.com/hadwin918/code-generate/cmd/code-generate/gen"
	"github.com/hadwin918/code-generate/cmd/code-generate/hook"
	"github.com/hadwin918/code-generate/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var verbose = false

var cmdRoot = &cobra.Command{
	Use:     "code-generate",
	Version: version.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
	},
}

func init() {
	cmdRoot.PersistentFlags().BoolVarP(&verbose, "verbose", "v", verbose, "")

	cmdRoot.AddCommand(gen.CmdGen)
	cmdRoot.AddCommand(hook.CmdHook)
}

func main() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
