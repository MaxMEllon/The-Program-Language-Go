package github

type User struct {
	Id      int    `json:"id"`
	Login   string `json:"login"`
	Avatar  string `json:"avatar_url"`
	Type    string `json:"type"`
	HtmlURL string `json:"html_url"`
}
