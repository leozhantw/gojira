package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMux    *http.ServeMux
	testServer *httptest.Server

	testClient *Client
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testClient, _ = NewClient(testServer.URL, "test-token")
}

func teardown() {
	testServer.Close()
}

func TestClient_NewRequest(t *testing.T) {
	setup()
	defer teardown()

	urlStr := "/rest/api/2/issue/"
	body := &Issue{Key: "JIRA-100"}
	req, _ := testClient.NewRequest("GET", urlStr, body)

	t.Run("relative URL was expanded", func(t *testing.T) {
		want := testServer.URL + urlStr
		got := req.URL.String()
		assert.Equal(t, want, got)
	})

	t.Run("body was JSON encoded", func(t *testing.T) {
		want := `{"key":"JIRA-100","fields":null}` + "\n"
		body, _ := ioutil.ReadAll(req.Body)
		got := string(body)
		assert.Equal(t, want, got)
	})
}

func TestClient_Do(t *testing.T) {
	setup()
	defer teardown()

	type foo struct {
		A string
	}

	testMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		fmt.Fprint(w, `{"A":"a"}`)
	})

	req, _ := testClient.NewRequest("GET", "/", nil)
	got := new(foo)
	testClient.Do(req, got)

	want := &foo{"a"}
	assert.Equal(t, want, got)
}
