// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import "fmt"

//!+
import "crypto/sha256"

// pc[i] is the population count of i.
var pc [256]byte

func Init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func ShaCount(ptr1 *[32]byte, ptr2 *[32]byte) int {
	c := 0
	for i := 0; i < 32; i++ {
		c += int(pc[ptr1[i]^ptr2[i]])
	}
	return c
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	Init()
	fmt.Printf("%x\n%x\n%t\n%T\n%d\n", c1, c2, c1 == c2, c1, ShaCount(&c1, &c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
