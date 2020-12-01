package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/leozhantw/gojira/internal/rendor"
	"github.com/leozhantw/gojira/pkg/jira"
)

var issueListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		jql := "assignee = currentuser()"

		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		if project != "" {
			jql += fmt.Sprintf(" AND project = '%s'", project)
		}

		includeStatus, err := cmd.Flags().GetStringSlice("include")
		if err != nil {
			return err
		}
		if len(includeStatus) > 0 {
			jql += fmt.Sprintf(" AND status IN ('%s')", strings.Join(includeStatus, "','"))
		}

		excludeStatus, err := cmd.Flags().GetStringSlice("exclude")
		if err != nil {
			return err
		}
		if len(excludeStatus) > 0 {
			jql += fmt.Sprintf(" AND status NOT IN ('%s')", strings.Join(excludeStatus, "','"))
		}

		client, err := jira.NewClient(cfg.SiteURL, cfg.Token)
		if err != nil {
			return err
		}
		issues, _, err := client.Issue.Search(jql)
		if err != nil {
			return err
		}
		rendor.New(os.Stdout).IssueList(issues)
		return nil
	},
}

func init() {
	issueListCmd.Flags().StringP("project", "p", "", "specify a project")
	issueListCmd.Flags().StringSliceP("include", "i", nil, "include specified status")
	issueListCmd.Flags().StringSliceP("exclude", "e", []string{"done", "invalid", "close", "closed"}, "exclude specified status")
}
