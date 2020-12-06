/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads
numbers from its command-line arguments or from the standard input if there are no arguments,
and converts each number into units like temperature in Celsius and Fahrenheit,
length in feet and meters, weight in pounds and kilograms, and the like.

Sample output:
go run ex2-2.go 1
1°F = -17.22222222222222°C, 1°C = 33.8°F
1m = 0.3048ft, 1ft = 0.30478512648582745m
1lb = 0.9999988107kg, 1kg = 2.20462lb
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv0"
	"weightconv"
	"lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv0.Fahrenheit(t)
		c := tempconv0.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv0.FToC(f), c, tempconv0.CToF(c))
		
		m := lengthconv.Meters(t)
		ft := lengthconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToF(m), ft, lengthconv.FToM(ft))

		p := weightconv.Pounds(t)
		k := weightconv.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}
