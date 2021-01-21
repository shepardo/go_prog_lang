package main

import (
	"fmt"
	"popcount2"
)

func main() {
	//fmt.Printf("Hello World\n");
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount2.PopCount(uint64(i)));
	}
}