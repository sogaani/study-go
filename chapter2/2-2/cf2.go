// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"./lengthconv"
	"./tempconv"
	"./weightconv"
)

func convertTemp(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Printf("input     \tFahrenheit\tCelsius    \tKelvin\n")
	fmt.Printf("Fahrenheit\t%s\t%s\t%s\n",
		f, tempconv.FToC(f), tempconv.FToK(f))
	fmt.Printf("Celsius   \t%s\t%s\t%s\n",
		tempconv.CToF(c), c, tempconv.CToK(c))
	fmt.Printf("Kelvin    \t%s\t%s\t%s\n",
		tempconv.KToF(k), tempconv.KToC(k), k)
}

func convertWeight(t float64) {
	k := weightconv.Kilogram(t)
	p := weightconv.Pound(t)
	fmt.Printf("input   \tKilogram\tPound\n")
	fmt.Printf("Kilogram\t%s\t%s\n",
		k, weightconv.KToP(k))
	fmt.Printf("Pound   \t%s\t%s\n",
		weightconv.PToK(p), p)
}

func convertLength(t float64) {
	m := lengthconv.Meter(t)
	f := lengthconv.Feet(t)
	fmt.Printf("input\tMeter     \tFeet\n")
	fmt.Printf("Meter\t%s\t%s\n",
		m, lengthconv.MToF(m))
	fmt.Printf("Feet \t%s\t%s\n",
		lengthconv.FToM(f), f)
}

func main() {
	switch os.Args[1] {
	case "temp":
		for _, arg := range os.Args[2:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			convertTemp(t)
		}
	case "weight":
		for _, arg := range os.Args[2:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			convertWeight(t)
		}
	case "length":
		for _, arg := range os.Args[2:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			convertLength(t)
		}
	}
}

//!-
