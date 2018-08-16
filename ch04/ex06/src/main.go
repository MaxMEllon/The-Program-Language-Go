package main

import (
	"fmt"
)

func compress(target string) string {
	b := []byte(target)
	l := len(b)
	r := []byte{}
	for i := 0; i < l; i++ {
		switch b[i] {
		case 9, 10, 11, 12, 13, 32, 133, 160:
			r = append(r, '.')
		default:
			r = append(r, b[i])
		}

	}
	return string(r)
}

func main() {
	// fmt.Println("blank characters")
	// fmt.Println(string([]byte{9}))
	// fmt.Println(string([]byte{10}))
	// fmt.Println(string([]byte{11}))
	// fmt.Println(string([]byte{12}))
	// fmt.Println(string([]byte{13}))
	// fmt.Println(string([]byte{32}))
	// fmt.Println(string([]byte{133}))
	// fmt.Println(string([]byte{160}))
	fmt.Println(compress(`aaa u             intptr			 
	
	
	
	
	`))
}
