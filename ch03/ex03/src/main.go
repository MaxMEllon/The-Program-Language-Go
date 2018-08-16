package main

import (
	"fmt"
	"math"
)

const (
	width, height = 500, 300
	cells         = 100
	xyrange       = 47.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			z := int64(z(i, j) * 470)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgb(%d, 0, %d);'/>\n", ax, ay, bx, by, cx, cy, dx, dy, z, 256-z)
		}
	}
	fmt.Println("</svg>")
}

func z(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return f(x, y)
}

func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, fmt.Errorf("Infinity")
	}

	if math.IsNaN(z) {
		return 0, 0, fmt.Errorf("Not a Number")
	}

	sx := width/2 + (x-y)*cos30*xyscale + 40
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	z := math.Sqrt(math.Sin(y/4)-math.Cos(x/4)) / 2
	if math.IsNaN(z) {
		z = 0.0
	}
	return z
}
