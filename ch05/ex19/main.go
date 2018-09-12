package main

import (
	"fmt"
)

func printMessage(s string) (r string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoverd")
			r = err.(string)
		}
	}()
	panic(s)
}

func main() {
	fmt.Printf("return: %+v\n", printMessage("any message"))
}
