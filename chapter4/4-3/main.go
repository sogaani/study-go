// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [256]int{0, 1, 2, 3, 4, 5}
	reverseArray(&a, 6)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!+rev
// reverse reverses a slice of ints in place.
func reverseArray(s *[256]int, len int) {
	for i, j := 0, len-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev