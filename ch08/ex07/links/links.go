// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 138.
//!+Extract

// Package links provides a link-extraction function.
package links

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	visitNode := func(n *html.Node, dir string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				_, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					fmt.Print(a.Val)
					continue // ignore bad URLs
				}
				exe, err := os.Executable()
				base := filepath.Dir(exe)
				var fp string
				val := strings.Replace(a.Val, "http://", "", -1)
				val = strings.Replace(val, "https://", "", -1)
				if strings.HasPrefix(val, "/") {
					fp = base + "/" + dir + val + ".html"
				} else {
					fp = base + "/" + dir + "/" + val + ".html"
				}
				n.Attr[i].Val = "file:///" + fp
				d := filepath.Dir(fp)

				os.MkdirAll(d, os.ModePerm)
				if a.Val != "/" && a.Val != "#" && !strings.Contains(fp, "#") {
					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
					}
					resp, err := http.Get(link.String())
					if !strings.HasPrefix(resp.Status, "200") {
						fmt.Printf("%s: %s\n", "SKIP", a.Val)
						continue
					} else {
						fmt.Printf("%s: %s\n", resp.Status, a.Val)
					}
					f, err := os.Create(fp)
					if err != nil {
						log.Fatal(err)
					}
					data, err := ioutil.ReadAll(resp.Body)
					f.Write(data)
					f.Close()

				}
			}
		}
	}
	f, err := os.Create(directory + "/index.html")
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	defer f.Close()
	forEachNode(doc, visitNode, nil, directory)
	html.Render(f, doc)
	return nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node, dir string), dir string) {
	if pre != nil {
		pre(n, dir)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, dir)
	}
	if post != nil {
		post(n, dir)
	}
}
