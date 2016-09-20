// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
	"unicode"
)

func delSpace(b []byte) []byte {
	i := 0
	for _, s := range b {
		if !unicode.IsSpace(rune(s)) || i == 0 || s != b[i-1] {
			b[i] = s
			i++
		}
	}
	return b[:i]
}

func main() {
	//!+main
	data := "a    b   c"
	fmt.Printf("%s\n", delSpace([]byte(data))) //a b c
	fmt.Printf("%s\n", data)                   //a    b   c

	//!-main
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

//!-alt
