package main

import (
	"flag"
	"fmt"
	"github.com/maxmellon/The-Program-Language-Go/ch05/ex13/links"
	"sync"
	"log"
	"math"
)

type leveledList struct {
	depth    int
	linkList []string
}

var depthOption *int
var n int

func init() {
	depthOption = flag.Int("depth", math.MaxInt32, "depth of links")
	flag.Parse()
	n = 0
}
var wg sync.WaitGroup

func crawl(depth int, url string) *leveledList {
	if depth > *depthOption {
		return &leveledList{depth + 1, make([]string, 0)}
		wg.Done()
		n--
	}

	fmt.Printf("depth %5d: url: %s\n", depth, url)

	list, err := links.Extract(url)
	wg.Done()
	n--

	if err != nil {
		log.Print(err)
	}

	return &leveledList{depth + 1, list}
}

func main() {
	workList := make(chan *leveledList)
	go func() {
		workList <- &leveledList{1, flag.Args()}
	}()


	seen := make(map[string]bool)

	for {
		list := <-workList
		for _, link := range list.linkList {
			if !seen[link] {
				seen[link] = true
				n++
				wg.Add(1)
				go func(depth int, link string) {
					workList <- crawl(depth, link)
				}(list.depth, link)
			}
			fmt.Println(n)
		}
		wg.Wait()
		if n == 0 && list.depth > *depthOption {
			break
		}
	}
}