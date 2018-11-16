package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		bytes   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			[]byte("a"),
			map[rune]int{'a': 1},
			[]int{0},
			0,
		},
		{
			[]byte("あ"),
			map[rune]int{'あ': 1},
			[]int{0},
			0,
		},
		{
			[]byte("👨‍👩‍👧"),
			map[rune]int{'👨': 1, '👩': 1, '👧': 1, '\u200d': 2},
			[]int{0},
			0,
		},
	}
	for _, testCase := range tests {
		counts, utflen, _, err := charcount(bytes.NewReader(testCase.bytes))

		if err != nil {
			t.Errorf("%v\n", err)
		}

		for idx, t := range testCase.counts {
			count := counts[idx]

			if count != t {
				panic(fmt.Sprintf("count is %d, want %d\n", count, t))
			}

		}

		for idx, t := range testCase.utflen {
			act := utflen[idx]

			if act != t {
				panic(fmt.Sprintf("utflen is %d, want %d\n", t, act))
			}
		}
	}
}
