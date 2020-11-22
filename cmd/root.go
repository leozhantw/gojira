package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:           "gojira",
		Short:         "Save your time on jira operations",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmds = []*cobra.Command{
		configureCmd,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cmds...)
}
