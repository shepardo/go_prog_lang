package main

import "fmt"

func main() {
	fmt.Printf("gcd(2, 256) = %d\n", gcd(2, 256))
	fmt.Printf("gcd(11, 33) = %d\n", gcd(11, 33))
}

func gcd(x, y int) int {
  for y != 0 {
    x, y = y, x%y
  }
  return x
}