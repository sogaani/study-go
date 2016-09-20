// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of word

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}
	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

//!-
