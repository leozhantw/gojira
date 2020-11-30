package jira

import (
	"net/http"
	"net/url"
)

type IssueService struct {
	client *Client
}

type Issue struct {
	Key    string       `json:"key"`
	Fields *IssueFields `json:"fields"`
}

type IssueFields struct {
	Status  *Status `json:"status"`
	Summary string  `json:"summary"`
}

type Status struct {
	Name string `json:"name"`
}

type searchResult struct {
	Issues []Issue `json:"issues"`
}

func (s *IssueService) Search(jql string) ([]Issue, *http.Response, error) {
	u := url.URL{
		Path: "/rest/api/3/search",
	}
	uv := url.Values{}
	if jql != "" {
		uv.Add("jql", jql)
	}

	u.RawQuery = uv.Encode()

	req, err := s.client.NewRequest("GET", u.String(), nil)
	if err != nil {
		return []Issue{}, nil, err
	}

	v := new(searchResult)
	resp, err := s.client.Do(req, v)
	if err != nil {
		err = NewJiraError(resp, err)
	}
	return v.Issues, resp, err
}
