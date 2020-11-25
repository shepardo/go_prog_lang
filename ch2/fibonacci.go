package main

import "fmt"

func main() {
	fmt.Printf("fibonacci 10: %d\n", fib(10))
	fmt.Printf("fibonacci 0: %d\n", fib(0))
	fmt.Printf("fibonacci 1: %d\n", fib(1))
	fmt.Printf("fibonacci 2: %d\n", fib(2))
	fmt.Printf("fibonacci 5: %d\n", fib(5))
}

func fib(n int) int {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    x, y = y, x+y
  }
  return x
}