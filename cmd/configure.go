package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Set up your Jira Configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		var org string
		fmt.Print("Organization: ")
		if _, err := fmt.Scan(&org); err != nil {
			return err
		}
		viper.Set("organization", org)

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		filename := path.Join(home, ".gojira.yaml")
		return viper.WriteConfigAs(filename)
	},
}
