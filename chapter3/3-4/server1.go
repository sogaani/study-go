// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
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

var width, height = 600, 320                        // canvas size in pixels
var cells = 100                                     // number of grid cells
var xyrange = 30.0                                  // axis ranges (-xyrange..+xyrange)
var xyscale = float64(width) / 2 / float64(xyrange) // pixels per x or y unit
var zscale = float64(height) * 0.4                  // pixels per z unit
var angle = math.Pi / 6                             // angle of x, y axes (=30°)
var color = "#00ff00"
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func graph(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j)
			bx, by, berr := corner(i, j)
			cx, cy, cerr := corner(i, j+1)
			dx, dy, derr := corner(i+1, j+1)
			if !(aerr || berr || cerr || derr) {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)

			}
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z, err := f(x, y)

	if err {
		return 0, 0, true
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
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
		w.Header().Set("Content-Type", "image/svg+xml")

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			var err error
			switch k {
			case "width":
				width, err = strconv.Atoi(v[0])
				if err != nil {
					width = 600
				}
			case "height":
				height, err = strconv.Atoi(v[0])
				if err != nil {
					height = 320
				}
			case "cells":
				cells, err = strconv.Atoi(v[0])
				if err != nil {
					cells = 100
				}
			case "xyrange":
				xyrange, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					xyrange = 30.0
				}
			case "color":
				color = v[0]

			}
		}
		xyscale = float64(width) / 2 / float64(xyrange) // pixels per x or y unit
		zscale = float64(height) * 0.4                  // pixels per z unit
		graph(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
	//!+main
}

//!-main
