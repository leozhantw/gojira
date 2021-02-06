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
		var siteURL string
		fmt.Print("Your website URL (e.g. https://mysite.atlassian.net): ")
		if _, err := fmt.Scan(&siteURL); err != nil {
			return err
		}
		viper.Set("siteurl", siteURL)

		var account string
		fmt.Print("Your account (e.g. me@example.com): ")
		if _, err := fmt.Scan(&account); err != nil {
			return err
		}
		viper.Set("account", account)

		var token string
		fmt.Print("Your API token: ")
		if _, err := fmt.Scan(&token); err != nil {
			return err
		}
		viper.Set("token", token)

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		filename := path.Join(home, ".gojira.yaml")
		return viper.WriteConfigAs(filename)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
