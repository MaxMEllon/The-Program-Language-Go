package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// refs: https://www.codewhoop.com/tutorial/array-rotation-in-place
func rotate(data *[10]int, r int) {
	d := -1
	n := len(data)
	g := gcd(n, r)
	for i := 0; i < g; i++ {
		tmp := data[i]
		j := i
		for {
			d = (j + r) % n
			if d == i {
				break
			}
			data[j] = data[d]
			j = d
		}
		data[j] = tmp
	}
}

func main() {
	data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(data)

	rotate(&data, 3)
	fmt.Println(data)
}
