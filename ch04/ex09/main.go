package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	i := bufio.NewScanner(os.Stdin)
	i.Split(bufio.ScanWords)
	for i.Scan() {
		counts[i.Text()]++
	}
	for text, n := range counts {
		fmt.Printf("%s\t%d\n", text, n)
	}
}
