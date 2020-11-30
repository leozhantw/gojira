package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/leozhantw/gojira/internal/rendor"
	"github.com/leozhantw/gojira/pkg/jira"
)

var issueListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := jira.NewClient(cfg.SiteURL, cfg.Token)
		if err != nil {
			return err
		}
		issues, _, err := client.Issue.Search("assignee = currentuser() AND status != done AND status != invalid AND status != close AND status != closed")
		if err != nil {
			return err
		}
		rendor.New(os.Stdout).IssueList(issues)
		return nil
	},
}
