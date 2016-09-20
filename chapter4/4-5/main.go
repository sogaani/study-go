// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

func delDuplicate(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != strings[i] {
			strings[i+1] = s
			i++
		}
	}
	return strings[:i+1]
}

func main() {
	//!+main
	data := []string{"one", "three", "three", "three", "one", "two", "two"}
	fmt.Printf("%q\n", delDuplicate(data)) //["one" "three" "one" "two"]
	fmt.Printf("%q\n", data)               //["one" "three" "one" "two" "one" "two" "two"]
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
