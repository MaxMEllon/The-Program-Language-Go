package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/maxmellon/The-Program-Language-Go/ch05/ex13/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func savePageFromURL(maybeUrl string) (err error) {
	url, err := url.Parse(maybeUrl)
	if err != nil {
		return
	}

	directoryPath := filepath.Join("downloads", url.Host, url.Path)
	if err = os.MkdirAll(directoryPath, 0755); err != nil {
		return
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	filePath := filepath.Join(directoryPath, "dump.html")
	wBuf, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer wBuf.Close()
	_, err = io.Copy(wBuf, resp.Body)
	if err != nil {
		return
	}
	return nil
}

func crawl(url string) []string {
	fmt.Println(url)
	savePageFromURL(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
