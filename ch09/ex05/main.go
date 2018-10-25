package main

import (
	"flag"
	"fmt"
	"time"
)

var pinch chan struct{}
var ponch chan struct{}
var done chan struct{}

var max *int
var count int

func pin() {
	<-ponch
	go pin()
	pinch <- struct{}{}
}

func pon() {
	<-pinch
	count++
	go pon()
	ponch <- struct{}{}
}

func init() {
	max = flag.Int("max", 1000000, "maximum call count")
	flag.Parse()
	count = 0
}

func main() {
	pinch = make(chan struct{})
	ponch = make(chan struct{})
	done = make(chan struct{})
	start := time.Now()
	go pin()
	go pon()
	time.AfterFunc(1*time.Second, func() {
		done <- struct{}{}
	})
	ponch <- struct{}{}
	<-done
	diff := time.Now().Sub(start)
	fmt.Printf("diff: %f seconds\n", float64(float64(diff)/float64(1000*1000*1000)))
	fmt.Printf("count: %d rally\n", count)
}
