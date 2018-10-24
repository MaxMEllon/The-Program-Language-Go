// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
//!+Extract

// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	host := resp.Request.URL.Host

	if err := os.Mkdir(host, os.ModePerm); err != nil {
		if os.IsExist(err) {
			fmt.Printf("%s directory exists! Please delete it\n", host)
			return nil
		}

		fmt.Printf("os.Mkdir : %v\n", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	extractAsFile(host, resp)
	defer resp.Body.Close()

	return nil
}

func extractAsFile(directory string, resp *http.Response) error {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				_, err := resp.Request.URL.Parse(a.Val)
				exe, err := os.Executable()
				dir := filepath.Dir(exe)
				n.Attr[i].Val = "file:" + dir + "/" + directory + a.Val + ".html"
				if err != nil {
					continue // ignore bad URLs
				}
				if a.Val != "/" && a.Val != "#" {
					_, err := http.Get(a.Val)
					if err != nil {

					}
					_, err = os.Create( dir + "/" + directory + a.Val + ".html")
					if err != nil {

					}
				}
				fmt.Println(n.Attr[i].Val)
			}
		}
	}
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	defer f.Close()
	html.Render(f, doc)
	forEachNode(doc, visitNode, nil)
	return nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
