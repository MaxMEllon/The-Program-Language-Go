// refs: http://stdkmd.com/blog/2017/04/newtons-method-and-fractal.html

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	x, err := strconv.ParseFloat(r.Form.Get("x"), 64)
	if err != nil {
		log.Print(err)
		return
	}

	y, err := strconv.ParseFloat(r.Form.Get("y"), 64)
	if err != nil {
		log.Print(err)
		return
	}

	render(w, x, y)
	return
}

func render(w io.Writer, argX, argY float64) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin + argX
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin + argY
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(w, img)
}

func newton(z complex128) color.Color {
	const iterations = 255
	const contrast = 1

	v := z
	for i := uint8(0); i < iterations; i++ {
		v -= ((v*v*v*v - 1) / (4 * v * v * v))
		h := 255 - contrast*i
		switch {
		case cmplx.Abs(1-v) < 1e-5:
			return color.RGBA{h, 0, 0, 255}
		case cmplx.Abs(-1-v) < 1e-5:
			return color.RGBA{0, h, 0, 255}
		case cmplx.Abs(1i-v) < 1e-5:
			return color.RGBA{0, 0, h, 255}
		case cmplx.Abs(-1i-v) < 1e-5:
			return color.RGBA{h, h, 0, 255}
		}
	}
	return color.Black
}
