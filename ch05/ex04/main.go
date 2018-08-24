package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(attr map[[2]string]string, n *html.Node) map[[2]string]string {
	if n == nil {
		return attr
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			attr[[...]string{n.Data, a.Key}] = a.Val
		}
	}
	attr = visit(attr, n.FirstChild)
	return visit(attr, n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for attr, val := range visit(make(map[[2]string]string), doc) {
		fmt.Printf("<%s %s=\"%s\" />\n", attr[0], attr[1], val)
	}
}
