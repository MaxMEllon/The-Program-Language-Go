package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprintf(os.Stderr, strings.Join(os.Args[1:], " "))
}
