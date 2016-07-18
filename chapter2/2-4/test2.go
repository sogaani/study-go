// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"time"

	"./popcount"
)

func main() {
	start := time.Now()
	for i := 0; i < 1024*1024; i++ {
		popcount.PopCount_1(0x1234567890ABCDEF)
	}
	end := time.Now()
	fmt.Printf("テーブル版PopCount:%f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	for i := 0; i < 1024*1024; i++ {
		popcount.PopCount_2(0x1234567890ABCDEF)
	}
	end = time.Now()
	fmt.Printf("Loop版PopCount:%f秒\n", (end.Sub(start)).Seconds())

	start = time.Now()
	for i := 0; i < 1024*1024; i++ {
		popcount.PopCount_3(0x1234567890ABCDEF)
	}
	end = time.Now()
	fmt.Printf("ビットシフト版PopCount:%f秒\n", (end.Sub(start)).Seconds())
}

//!-
