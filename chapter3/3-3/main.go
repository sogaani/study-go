// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	maxz          = 1
	minz          = -2 / math.Pi
	unit          = (maxz - minz/4)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aerr := corner(i+1, j)
			bx, by, bz, berr := corner(i, j)
			cx, cy, cz, cerr := corner(i, j+1)
			dx, dy, dz, derr := corner(i+1, j+1)
			if !(aerr || berr || cerr || derr) {
				z := (az+bz+cz+dz)/4 - minz
				fill := "#0000ff"

				switch int(z / unit) {
				case 0:
					stride := int(z * 0xff / unit)
					s := fmt.Sprintf("%x", stride)
					fill = "#00" + s + "ff"
				case 1:
					stride := int(0xff - (z-unit)*0xff/unit)
					s := fmt.Sprintf("%x", stride)
					fill = "#00ff" + s
				case 2:
					stride := int((z - 2*unit) * 0xff / unit)
					s := fmt.Sprintf("%x", stride)
					fill = "#" + s + "ff00"
				case 3:
					stride := int(0xff - (z-2*unit)*0xff/unit)
					s := fmt.Sprintf("%x", stride)
					fill = "#ff" + s + "00"
				}

				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, fill)

			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, err := f(x, y)

	if err {
		return 0, 0, 0, true
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, false
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	ret := math.Sin(r) / r
	if math.IsInf(ret, 0) || math.IsNaN(ret) {
		return 0, true

	}
	return ret, false
}

//!-
