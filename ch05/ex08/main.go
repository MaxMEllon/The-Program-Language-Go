package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	elem := ElementByID(doc, os.Args[1])
	if elem != nil {
		printNode(elem)
	}
}

func printNode(n *html.Node) {
	fmt.Printf("<%s ", n.Data)
	for _, attr := range n.Attr {
		fmt.Printf("%s=\"%s\" ", attr.Key, attr.Val)
	}
	fmt.Printf("/>\n")
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	itr := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key != "id" || attr.Val != id {
					continue
				}
				node = n
				return false
			}
		}
		return true
	}
	forEachNode(doc, itr, itr)
	return node
}
