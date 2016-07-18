// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package weightconv

// Celsius Pound to Kilogram.
func PToK(p Pound) Kilogram { return Kilogram(p * ptok) }

// Convert Kilogram to Pound.
func KToP(k Kilogram) Pound { return Pound(k * ktop) }

//!-
