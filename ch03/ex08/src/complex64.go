// refs: http://stdkmd.com/blog/2017/04/newtons-method-and-fractal.html

package main

import (
	// "fmt"
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
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(complex64(z)))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z complex64) color.Color {
	const iterations = 255
	const contrast = 1

	v := z
	for i := uint8(0); i < iterations; i++ {
		v -= ((v*v*v*v - 1) / (4 * v * v * v))
		h := 255 - contrast*i
		switch {
		case cmplx.Abs(complex128(1-v)) < 1e-7:
			return color.RGBA{h, 0, 0, 255}
		case cmplx.Abs(complex128(-1-v)) < 1e-7:
			return color.RGBA{0, h, 0, 255}
		case cmplx.Abs(complex128(1i-v)) < 1e-7:
			return color.RGBA{0, 0, h, 255}
		case cmplx.Abs(complex128(-1i-v)) < 1e-7:
			return color.RGBA{h, h, 0, 255}
		}
	}
	return color.Black
}
