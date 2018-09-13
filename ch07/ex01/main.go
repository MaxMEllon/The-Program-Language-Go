package main

import (
	"fmt"
	"github.com/maxmellon/The-Program-Language-Go/ch07/ex01/counter"
)

func main() {
	var c1 counter.WordCounter
	var c2 counter.LineCounter

	str := `
		hello, wolrd 1 2
		hello, wolrd 1 2
		hello, wolrd 1 2
		hello, wolrd 1 2
		hello, wolrd 1 2
	`

	fmt.Fprintf(&c1, str)
	fmt.Fprintf(&c2, str)
	fmt.Println(c1)
	fmt.Println(c2)
}
