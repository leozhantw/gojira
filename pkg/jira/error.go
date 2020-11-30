package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

type Error struct {
	HTTPError     error
	ErrorMessages []string          `json:"errorMessages"`
	Errors        map[string]string `json:"errors"`
}

func NewJiraError(resp *http.Response, httpError error) error {
	if resp == nil {
		return errors.Wrap(httpError, "No response returned")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, httpError.Error())
	}
	jerr := Error{HTTPError: httpError}
	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.Unmarshal(body, &jerr)
		if err != nil {
			httpError = errors.Wrap(errors.New("could not parse JSON"), httpError.Error())
			return errors.Wrap(err, httpError.Error())
		}
	} else {
		if httpError == nil {
			return fmt.Errorf("got response status %s:%s", resp.Status, string(body))
		}
		return errors.Wrap(httpError, fmt.Sprintf("%s: %s", resp.Status, string(body)))
	}

	return &jerr
}

func (e *Error) Error() string {
	if len(e.ErrorMessages) > 0 {
		return fmt.Sprintf("%s: %v", e.ErrorMessages[0], e.HTTPError)
	}
	if len(e.Errors) > 0 {
		for key, value := range e.Errors {
			return fmt.Sprintf("%s - %s: %v", key, value, e.HTTPError)
		}
	}
	return e.HTTPError.Error()
}
