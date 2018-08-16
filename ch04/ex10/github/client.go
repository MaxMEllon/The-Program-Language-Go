package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(query []string) (*Issues, error) {
	q := url.QueryEscape(strings.Join(query, " "))

	resp, err := http.Get(SearchIssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %s", resp.Status)
	}

	var result Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
