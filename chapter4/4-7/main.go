// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) []byte {
	for i := len(b); i > 0; {
		_, size := utf8.DecodeRune(b)
		rotate(b[:i], size)
		i -= size
	}
	return b
}

func rotate(s []byte, n int) {
	for i, j := 0, n; i < len(s)-n; i, j = i+1, j+1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	//!+main
	data := "あ   い   　う"
	fmt.Printf("%s\n", reverse([]byte(data))) //a b c
	fmt.Printf("%s\n", data)                  //a    b   c

	//!-main
}
