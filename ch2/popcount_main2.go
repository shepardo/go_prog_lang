package main

import (
	"fmt"
	"popcount"
)

func main() {
	//fmt.Printf("Hello World\n");
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount.PopCount(uint64(i)));
	}
}