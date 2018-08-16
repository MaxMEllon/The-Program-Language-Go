package github

import "time"

type Issues struct {
	Items []*Issue
}

type Issue struct {
	Id            int       `json:"id"`
	Number        int       `json:"number"`
	RepositoryURL string    `json:"repository_url"`
	State         string    `json:"state"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	HtmlURL       string    `json:"html_url"`
	User          *User     `json:"user"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
