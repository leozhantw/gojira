package rendor

import (
	"io"

	"github.com/olekukonko/tablewriter"

	"github.com/leozhantw/gojira/pkg/jira"
)

type Rendor struct {
	out io.Writer
}

func New(out io.Writer) *Rendor {
	return &Rendor{
		out: out,
	}
}

func (r *Rendor) IssueList(issues []jira.Issue) {
	var data [][]string
	for _, issue := range issues {
		data = append(data, []string{
			issue.Key,
			issue.Fields.Summary,
			issue.Fields.Status.Name,
		})
	}

	table := tablewriter.NewWriter(r.out)
	table.SetHeader([]string{"key", "summary", "status"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
	)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiMagentaColor},
	)
	table.AppendBulk(data)
	table.Render()
}
