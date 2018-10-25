package main

import (
	"flag"
	"fmt"
	"time"
)

var pipelineNum *int

func init() {
	pipelineNum = flag.Int("p", 1000, "pipeline number")
	flag.Parse()
}

func main() {
	var start time.Time
	var end time.Time
	var diff time.Duration
	var sum int64
	done := make(chan int)
	for i := 0; i < *pipelineNum; i++ {
		next := make(chan int)
		prev := make(chan int)
		go func(i int) {
			start = time.Now()
			prev <- i
		}(i)
		go func() {
			v := <-prev
			next <- v
		}()
		go func() {
			v := <-next
			end = time.Now()
			diff = end.Sub(start)
			sum += diff.Nanoseconds()
			if v == *pipelineNum-1 {
				done <- v
			}
		}()
	}
	<-done
	fmt.Printf("sum: %20d nano second\n", sum)
	fmt.Printf("average: %20d nano second\n", sum/int64(*pipelineNum))
}
