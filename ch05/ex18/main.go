package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	fileClose := func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}
	defer fileClose()
	n, err = io.Copy(f, resp.Body) // fileクローズ前に呼ばれる
	return local, n, err
}

func main() {
	for _, s_url := range os.Args[1:] {
		filename, n, err := fetch(s_url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", s_url, err)
			continue
		}
		fmt.Printf("save %s, %d bytes\n", filename, n)
	}
}
