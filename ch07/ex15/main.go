package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

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
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	env := eval.Env{}
	if len(vars) != 0 {
		fmt.Printf("変数を入力\n")
	}
	for v := range vars {
		fmt.Printf("%s = ", v)
		stdin.Scan()
		valueStr := stdin.Text()
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		env[v] = value
	}
	fmt.Println(expr.Eval(env))
}
