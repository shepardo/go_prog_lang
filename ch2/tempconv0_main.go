package main

import (
	"fmt"
	"tempconv0"
)

func main() {
	fmt.Printf("%g\n", tempconv0.BoilingC - tempconv0.FreezingC)	// "100" °C
	boilingF := tempconv0.CToF(tempconv0.BoilingC)
	fmt.Printf("%g\n", boilingF - tempconv0.CToF(tempconv0.FreezingC)) // "180" °F
	//fmt.Printf("%g\n", boilingF - tempconv0.FreezingC)	// compile error: type mismatch

	var c tempconv0.Celsius
	var f tempconv0.Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == tempconv0.Celsius(f)) // "true"!

	c = tempconv0.FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c) // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c) // "100°C"
	fmt.Println(c) // "100°C"
	fmt.Printf("%g\n", c) // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}