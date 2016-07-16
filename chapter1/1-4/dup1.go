// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	mapfiles := make(map[string][]string)
	fnum := len(files)
	if fnum == 0 {
		countLines(os.Stdin, counts, 0, mapfiles, 1)
	} else {
		for i, arg := range files {

			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, i, mapfiles, fnum)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d回出現:\t%s\t出現するファイル:%s\n", n, line, strings.Join(mapfiles[line], "\t"))
		}
	}
}

func countLines(f *os.File, counts map[string]int, i int, mapfiles map[string][]string, fnum int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if mapfiles[line] == nil {
			mapfiles[line] = make([]string, fnum)
		}
		mapfiles[line][i] = f.Name()
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
