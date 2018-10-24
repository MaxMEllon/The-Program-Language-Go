package main

import (
	"fmt"
	"github.com/maxmellon/The-Program-Language-Go/ch08/ex10/links"
	"log"
	"os"
)

var n int

func init() {
	n = 0
}

var ch = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Printf("url: %s\n", url)
	ch <- struct{}{}
	list, err := links.Extract(url)
	<-ch
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	go func() {
		os.Stdin.Read(make([]byte, 255))
		links.Close()
	}()

	n++
	go func() {
		workList <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
