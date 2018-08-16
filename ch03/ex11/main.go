package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteString(s[:1])
		s = s[1:]
	}

	tail := ""
	if idx := strings.IndexByte(s, '.'); idx >= 0 {
		s, tail = s[:idx], s[idx:]
	}

	n := len(s)
	for i, v := range s {
		buf.Write([]byte(string(v)))
		if i%3 == 0 && n-1 != i {
			buf.WriteString(",")
		}
	}
	if len(tail) > 0 {
		buf.WriteString(tail)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma(os.Args[1]))
}
