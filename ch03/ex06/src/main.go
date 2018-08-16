package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var dy, dx float64
	dy, dx = (height/2)*(ymax-ymin), (width/2)*(xmax-xmin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			list := make([]color.Color, 0)
			list = append(list, mandelbrot(complex(x-dx, y)))
			list = append(list, mandelbrot(complex(x, y-dy)))
			list = append(list, mandelbrot(complex(x+dx, y)))
			list = append(list, mandelbrot(complex(x, y+dy)))
			img.Set(px, py, averageColor(list))
		}
	}
	png.Encode(os.Stdout, img)
}

func averageColor(list []color.Color) color.Color {
	r, g, b := 0, 0, 0
	for _, c := range list {
		dr, dg, db, _ := c.RGBA()
		r += int(dr)
		g += int(dg)
		b += int(db)
	}
	x := len(list)
	return color.RGBA{uint8(r / x), uint8(g / x), uint8(b / x), 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 1

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			h := 255 - contrast*n
			return color.YCbCr{h, uint8(cmplx.Abs(v)), uint8(cmplx.Abs(v))}
		}
	}
	return color.Black
}
