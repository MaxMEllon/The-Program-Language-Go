package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
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
	}
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		log.Print(err)
	}
	lissajous(w, cycles)

}

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x88, 0x00, 0xFF},
	color.RGBA{0x00, 0x66, 0x00, 0xFF},
	color.RGBA{0x00, 0x44, 0x00, 0xFF},
	color.RGBA{0x22, 0x22, 0x00, 0xFF},
	color.RGBA{0x44, 0x00, 0x00, 0xFF},
	color.RGBA{0x88, 0x00, 0x00, 0xFF},
	color.RGBA{0xAA, 0x00, 0x00, 0xFF},
	color.RGBA{0xCC, 0x00, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0xCC, 0x00, 0xCC, 0xFF},
	color.RGBA{0xAA, 0x00, 0xAA, 0xFF},
	color.RGBA{0x88, 0x00, 0x88, 0xFF},
	color.RGBA{0x66, 0x00, 0x66, 0xFF},
	color.RGBA{0x44, 0x00, 0x44, 0xFF},
	color.RGBA{0x22, 0x22, 0x22, 0xFF},
	color.RGBA{0x00, 0x44, 0x44, 0xFF},
	color.RGBA{0x00, 0x66, 0x66, 0xFF},
	color.RGBA{0x00, 0x88, 0x88, 0xFF},
	color.RGBA{0x00, 0xAA, 0xAA, 0xFF},
	color.RGBA{0x00, 0xCC, 0xCC, 0xFF},
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 200
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	index := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		if i == 0 || (i%len(palette)-1) == 0 {
			index = 1
		}
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(index))
		}
		index++
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
