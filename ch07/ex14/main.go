package main

import (
	"fmt"

	"github.com/maxmellon/The-Program-Language-Go/ch07/ex13/eval"
)

func main() {
	expr, _ := eval.Parse("pow(x, 3) + pow(y, 3)")
	fmt.Println(expr.String())
}
