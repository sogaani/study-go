// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	ptok Pound    = Pound(1 / ktop)
	ktop Kilogram = 0.45359237
)

func (p Pound) String() string    { return fmt.Sprintf("%10.3flb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%10.3fKg", k) }

//!-
