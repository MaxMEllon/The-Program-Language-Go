package main

import "fmt"

func join(sep string, strings ...string) (result string) {
	result = ""
	length := len(strings)
	for idx, s := range strings {
		if length-1 == idx {
			result += s
			break
		}
		result += s + sep
	}
	return
}

func main() {
	fmt.Println(join(", ", "hoge", "poge", "foo", "bar"))
}
