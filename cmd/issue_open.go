package cmd

import (
	"fmt"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var issueOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open by your default browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("%s/browse/%s", cfg.SiteURL, args[0])
		return browser.OpenURL(url)
	},
}
