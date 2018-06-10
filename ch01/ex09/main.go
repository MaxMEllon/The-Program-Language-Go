package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func isURL(maybeURL string) bool {
	r := regexp.MustCompile(`^http://*`)
	return r.Match([]byte(maybeURL))
}

func main() {
	for _, url := range os.Args[1:] {
		if !isURL(url) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status: %s\n", resp.Status)

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
