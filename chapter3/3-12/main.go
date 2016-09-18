// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("  %t\n", isAnagrams(os.Args[1], os.Args[2]))
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func isAnagrams(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	for i := 0; i < n; i++ {
		if strings.Count(s1, string(s1[i])) != strings.Count(s2, string(s1[i])) {
			return false
		}
	}
	return true
}

//!-
