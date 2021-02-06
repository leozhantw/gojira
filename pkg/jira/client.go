package jira

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	client *http.Client

	baseURL *url.URL
	account string
	token   string

	Issue *IssueService
}

func NewClient(baseURL, account, token string) (*Client, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client:  http.DefaultClient,
		baseURL: parsedBaseURL,
		account: account,
		token:   token,
	}
	c.Issue = &IssueService{client: c}
	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// Relative URLs should be specified without a preceding slash since baseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(context.Background(), method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	str := fmt.Sprintf("%s:%s", "leozhan@kkstream.com", c.token)
	encStr := base64.StdEncoding.EncodeToString([]byte(str))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encStr))

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c := resp.StatusCode; c < 200 || c > 299 {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, fmt.Errorf("request failed. Please analyze the request body for more details. Status code: %d", c)
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
