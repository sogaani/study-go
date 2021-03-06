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
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := strings.Index(s, ".")
	if n < 0 {
		n = len(s)
	}

	var b bytes.Buffer
	var i = n % 3
	if i == 0 || (i == 1 && strings.Index(s, "-") == 0) {
		i += 3
	}
	b.WriteString(s[:i])

	for i += 3; i <= n; i += 3 {
		b.WriteString(",")
		b.WriteString(s[i-3 : i])
	}
	b.WriteString(s[n:len(s)])
	return b.String()
}

//!-
