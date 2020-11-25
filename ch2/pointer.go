package main

import "fmt"

func main() {
	p := new(int) // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	*p = 2 // sets the unnamed int to 2
	fmt.Println(*p) // "2"
}

func newInt() *int { 
	return new(int) 
}

func newInt2() *int {	
	var dummy int
	return &dummy
}