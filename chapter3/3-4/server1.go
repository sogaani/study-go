// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
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
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j)
			bx, by, berr := corner(i, j)
			cx, cy, cerr := corner(i, j+1)
			dx, dy, derr := corner(i+1, j+1)
			if !(aerr || berr || cerr || derr) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)

			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, err := f(x, y)

	if err {
		return 0, 0, true
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, false
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	ret := math.Sin(r) / r
	if math.IsInf(ret, 0) || math.IsNaN(ret) {
		return 0, true

	}
	return ret, false
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles := 5   // number of complete x oscillator revolutions
		res := 0.001  // angular resolution
		size := 100   // image canvas covers [-size..+size]
		nframes := 64 // number of animation frames
		delay := 8    // delay between frames in 10ms units
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			var err error
			switch k {
			case "cycles":
				cycles, err = strconv.Atoi(v[0])
				if err != nil {
					cycles = 5
				}
			case "res":
				res, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					res = 0.001
				}
			case "size":
				size, err = strconv.Atoi(v[0])
				if err != nil {
					size = 100
				}
			case "nframes":
				nframes, err = strconv.Atoi(v[0])
				if err != nil {
					nframes = 64
				}
			case "delay":
				delay, err = strconv.Atoi(v[0])
				if err != nil {
					delay = 8
				}
			}
		}
		lissajous(w, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
	//!+main
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
