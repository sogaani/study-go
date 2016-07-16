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
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main
var colornum int = 256
var palette []color.Color

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

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func makePalette() {
	palette = make([]color.Color, colornum)
	stride := colornum / 3
	//赤から青
	var r uint8 = 255
	var g uint8 = 0
	var b uint8 = 0
	var diff uint8 = uint8(256 / stride)
	i := 0
	for ; i < stride; i++ {
		palette[i] = color.RGBA{r, g, b, 0xff}
		r -= diff
		b += diff
	}
	//青から緑
	r = 0
	g = 0
	b = 255
	for ; i < stride*2; i++ {
		palette[i] = color.RGBA{r, g, b, 0xff}
		b -= diff
		g += diff
	}
	//緑から赤
	r = 0
	g = 255
	b = 0
	for ; i < stride*3; i++ {
		palette[i] = color.RGBA{r, g, b, 0xff}
		g -= diff
		r += diff
	}

	//あまりがあれば
	for ; i < colornum; i++ {
		palette[i] = color.RGBA{r, g, b, 0xff}
	}
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	diff := 2 * math.Pi / 256.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := 0.0
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
			colorIndex += diff
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
