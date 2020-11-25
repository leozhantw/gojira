package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open by your default browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("https://%s.atlassian.net/browse/%s", cfg.Organization, args[0])
		return browser.OpenURL(url)
	},
}
