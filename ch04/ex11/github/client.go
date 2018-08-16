package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	Url   string
	Token string
}

func NewClient(token string) (*Client, error) {
	u, err := url.ParseRequestURI(BaseURL)
	if err != nil {
		return nil, err
	}
	return &Client{u.String(), token}, nil
}

func (c *Client) get(url string, t interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetIssues(repo, user string) (*[]Issue, error) {
	var result *[]Issue
	err := get(c.Url+path.Join("repos", user, repo, "issues"), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) GetIssue(repo, user, number string) (*Issue, error) {
	var result *Issue
	err := get(c.Url+path.Join("repos", user, repo, "issues", number), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
