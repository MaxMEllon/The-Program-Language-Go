package main

import (
	"fmt"
	"os"

	"github.com/maxmellon/The-Program-Language-Go/ch07/ex02/counter"
)

func main() {
	c, v := counter.CountingWriter(os.Stdout)
	c.Write([]byte("hello, world"))
	fmt.Printf("\n0x%x\n", v)
}
