package rendor

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/leozhantw/gojira/pkg/jira"
)

func TestRendor_IssueList(t *testing.T) {
	buf := &bytes.Buffer{}
	issues := []jira.Issue{
		{
			Key: "JIRA-100",
			Fields: &jira.IssueFields{
				Status:  &jira.Status{Name: "To Do"},
				Summary: "Summary 001",
			},
		},
		{
			Key: "JIRA-101",
			Fields: &jira.IssueFields{
				Status:  &jira.Status{Name: "In Progress"},
				Summary: "Summary 002",
			},
		},
	}

	New(buf).IssueList(issues)

	contains := []string{
		"KEY", "SUMMARY", "STATUS",
		"JIRA-100", "Summary 001", "To Do",
		"JIRA-101", "Summary 002", "In Progress",
	}
	s := buf.String()
	for _, c := range contains {
		assert.Contains(t, s, c)
	}
}
