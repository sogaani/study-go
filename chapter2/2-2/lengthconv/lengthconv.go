// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
	mtof Meter = 3.2808
	ftom Feet  = Feet(1 / mtof)
)

func (c Feet) String() string  { return fmt.Sprintf("%10.3fFt.", c) }
func (f Meter) String() string { return fmt.Sprintf("%10.3fm", f) }

//!-
