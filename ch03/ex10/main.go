package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	for i, v := range s {
		buf.Write([]byte(string(v)))
		if i%3 == 0 && n-1 != i {
			buf.Write([]byte(","))
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("sajdkhgfjakhgdfkajsfgd"))
}
