// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	balckIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		x := 0
		y := 0
		res := 1024

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			var err error
			switch k {
			case "x":
				x, err = strconv.Atoi(v[0])
				if err != nil {
					x = 0
				}
			case "y":
				y, err = strconv.Atoi(v[0])
				if err != nil {
					y = 0
				}
			case "res":
				res, err = strconv.Atoi(v[0])
				if err != nil {
					res = 1024
				}
			}
		}
		flacta(w, x, y, res)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
	//!+main
}

func flacta(out io.Writer, tx int, ty int, res int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)

	img := image.NewRGBA(image.Rect(0, 0, res, res))
	for py := 0; py < res; py++ {
		y := float64(py)/float64(res)*(ymax-ymin) + ymin + float64(ty)
		for px := 0; px < res; px++ {
			x := float64(px)/float64(res)*(xmax-xmin) + xmin + float64(tx)
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - contrast*n, 255 - contrast*n, 255 - contrast*n}
		}
	}
	return color.Black
}

//!-main
