package cmd

import (
	"github.com/spf13/cobra"
)

var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
}

func init() {
	rootCmd.AddCommand(issueCmd)
	issueCmd.AddCommand([]*cobra.Command{
		issueListCmd,
		issueOpenCmd,
	}...)
}
