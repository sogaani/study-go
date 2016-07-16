// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end := time.Now()
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	start = time.Now()
        fmt.Println(strings.Join(os.Args[1:], " "))
	end = time.Now()
        fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
}

//!-
