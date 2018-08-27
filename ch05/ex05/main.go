package main

import (
	"net/http"

	"golang.org/x/net/html"
)

func countWordsAndImages(n *html.Node) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	word, images = countWordsAndImages(doc)
	return
}
