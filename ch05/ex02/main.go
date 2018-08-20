package main

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func countDOM(doms map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return doms
	}
	if n.Type == html.ElementNode {
		doms[n.Data]++
	}
	doms = countDOM(doms, n.FirstChild)
	return countDOM(doms, n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	doms := countDOM(make(map[string]int), doc)
	names := make([]string, 0, len(doms))
	for name := range doms {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, k := range names {
		fmt.Printf("%-18s: %4d\n", k, doms[k])
	}
}
