package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/maxmellon/The-Program-Language-Go/ch07/ex13/eval"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	expr, err := eval.Parse(text)
	if err != nil {
		errLog := fmt.Errorf("failure parse %s, detail: %v", text, err)
		log.Fatal(errLog)
		os.Exit(1)
	}
	env := eval.Env{}
	fmt.Println(expr.Eval(env))
}
