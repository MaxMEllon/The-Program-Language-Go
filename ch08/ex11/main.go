package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	b, err := mirrorQuery(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}

func mirrorQuery(urls []string) ([]byte, error) {
	responses := make(chan *http.Response)
	cancel := make(chan struct{})
	defer close(cancel)
	for _, url := range urls {
		go func() {
			resp, _ := fetchWithCancel("GET", url, nil, cancel)
			responses <- resp
		}()
	}
	res := <- responses
	cancel <- struct{}{} // どれか一つでもレスポンス入れば他全部キャンセル
	return ioutil.ReadAll(res.Body)
}

func fetchWithCancel(method, url string, body io.Reader, cancel <-chan struct{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	req.Cancel = cancel
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}