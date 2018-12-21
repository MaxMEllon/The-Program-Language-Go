package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [2]string{"hello", "world"}
	slice := []string{"hello", "world"}
	mp := map[string]string{
		"hello": "world",
	}
	fmt.Println(reflect.ValueOf(arr).CanAddr())
	fmt.Println(reflect.ValueOf(slice).CanAddr())
	fmt.Println(reflect.ValueOf(mp).CanAddr())
	fmt.Println(reflect.ValueOf(&arr).CanAddr())
	fmt.Println(reflect.ValueOf(&slice).CanAddr())
	fmt.Println(reflect.ValueOf(&mp).CanAddr())
	fmt.Println(reflect.ValueOf(1).CanAddr())
}
