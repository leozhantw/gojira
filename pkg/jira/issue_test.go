package jira

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIssueService_Search(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/rest/api/3/search", func(w http.ResponseWriter, r *http.Request) {
		t.Run("method", func(t *testing.T) {
			assert.Equal(t, "GET", r.Method)
		})
		t.Run("request URL", func(t *testing.T) {
			want := "/rest/api/3/search?jql=test-jql"
			got := r.URL.String()
			assert.Equal(t, got, want)
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{}`)
	})

	_, _, err := testClient.Issue.Search("test-jql")
	require.NoError(t, err)
}
