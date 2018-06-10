package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

func isURL(maybeURL string) bool {
	r := regexp.MustCompile(`^http://*`)
	return r.Match([]byte(maybeURL))
}

func fetchAll(urlList []string) {
	start := time.Now()
	ch := make(chan string)

	for _, url := range urlList {
		if !isURL(url) {
			url = "http://" + url
		}
		go fetch(url, ch)
	}

	for range urlList {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("white reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

func main() {
	fetchAll(os.Args[1:])
}
