package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func isScriptOrStyle(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "script" || n.Data == "style"
}

func getAllPlaneText(planeText []string, n *html.Node) []string {
	if n == nil || isScriptOrStyle(n) {
		return planeText
	}
	if n.Type == html.TextNode {
		data := strings.Replace(n.Data, "\n\n", "", -1)
		planeText = append(planeText, data)
	}
	planeText = getAllPlaneText(planeText, n.FirstChild)
	return getAllPlaneText(planeText, n.NextSibling)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range getAllPlaneText(make([]string, 0), doc) {
		fmt.Printf("%s", text)
	}
}
