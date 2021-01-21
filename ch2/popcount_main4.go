package main

import (
	"fmt"
	"popcount4"
)

func main() {
	//fmt.Printf("Hello World\n");
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount4.PopCount(uint64(i)));
	}
}