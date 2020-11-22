package cmd

import (
	"fmt"
	"os"

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

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gojira")
		viper.SetConfigType("yaml")

		viper.Set("organization", org)

		return viper.WriteConfig()
	},
}
