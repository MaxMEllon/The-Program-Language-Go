package counter

import (
	"unicode"
	"unicode/utf8"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	isWord := true
	width := 0
	for i := 0; i < len(p); i += width {
		var r rune
		r, width = utf8.DecodeRune(p[i:])
		if unicode.IsSpace(r) {
			isWord = false
		} else {
			if !isWord {
				isWord = true
				*c++
			}
		}
	}
	return len(p), nil
}
