// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package lengthconv

// Convert Meter to Feet.
func MToF(m Meter) Feet { return Feet(m * mtof) }

// Convert Feet to Meter.
func FToM(f Feet) Meter { return Meter(f * ftom) }

//!-
