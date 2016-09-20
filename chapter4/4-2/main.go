// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt" //!+
	"os"
)

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
	// (name, default, help)
	var f = flag.Bool("f", false, "enable SHA384")
	var s = flag.Bool("s", false, "enable SHA512")
	flag.Parse()
	Init()
	if *f {
		c := sha512.Sum384([]byte(os.Args[1]))
		fmt.Printf("%x\n\n", c)
	} else if *s {
		c := sha512.Sum512([]byte(os.Args[1]))
		fmt.Printf("%x\n\n", c)
	} else {
		c := sha256.Sum256([]byte(os.Args[1]))
		fmt.Printf("%x\n\n", c)
	}
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
