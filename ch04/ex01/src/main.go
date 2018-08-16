package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	vim := sha256.New().Sum([]byte("vim"))
	emacs := sha256.New().Sum([]byte("emacs"))
	fmt.Printf("vim: %v\nemacs: %v\n", vim, emacs)

	diff := 0
	for idx := range vim {
		diff += int(pc[vim[idx]^emacs[idx]])
	}

	fmt.Printf("diff: %v\n", diff)
}
