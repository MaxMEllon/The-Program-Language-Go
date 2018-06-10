package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/maxmellon/The-Program-Language-Go/ch02/ex02/src/distconv"
	"github.com/maxmellon/The-Program-Language-Go/ch02/ex02/src/tempconv"
	"github.com/maxmellon/The-Program-Language-Go/ch02/ex02/src/weightconv"
)

var (
	d *bool
	t *bool
	w *bool
)

func init() {
	d = flag.Bool("d", false, "distance")
	t = flag.Bool("t", false, "temperature")
	w = flag.Bool("w", false, "weight")
}

func main() {
	flag.Parse()
	converter()
}

func converter() {
	for _, val := range flag.Args() {
		val, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Parse error")
			os.Exit(1)
		}
		if *w {
			g := weightconv.Gram(val)
			p := weightconv.Pound(val)
			fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", g, weightconv.GToP(g), p, weightconv.PToG(p))
		}
		if *d {
			m := distconv.Metre(val)
			y := distconv.Yard(val)
			fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", m, distconv.MToY(m), y, distconv.YToM(y))
		}
		if *t {
			c := tempconv.Celsius(val)
			f := tempconv.Fahrenheit(val)
			fmt.Fprintf(os.Stdout, "%s = %s, %s = %s\n", c, tempconv.CToF(c), f, tempconv.FToC(f))
		}
	}
}
