package main

import (
	"github.com/maxmellon/The-Program-Language-Go/ch08/ex07/links"
	"log"
	"os"
)

func main() {
	if err := links.Extract(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
