// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// Celsius temperature to .
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + AbsoluteZeroC) }

// Convert Fahrenheit temperature to .
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin  { return CToK(FToC(f)) }

func KToC(k Kelvin) Celsius    { return Celsius(k) - AbsoluteZeroC }
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

//!-
