package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 1

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 1024 {
			h := 255 - contrast*n
			k := int64(math.Abs(math.Cos(float64(h)) * 3))
			if k == 0 {
				return color.RGBA{h, 0, 0, 255}
			}
			if k == 1 {
				return color.RGBA{0, h, 0, 255}
			}
			if k == 2 {
				return color.RGBA{0, 0, h, 255}
			}
		}
	}
	return color.Black
}
