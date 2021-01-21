package main

import (
	"fmt"
	"popcount"
	"popcount2"	
	"popcount3"
	"popcount4"
	"time"
)

func main() {
	start := time.Now()
	//result := 0
	var time1, time2, time3, time4 int64
	start = time.Now()
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount.PopCount(uint64(i)))
	}
	time1 = time.Since(start).Microseconds()

	fmt.Printf("popcount v2\n");
	start = time.Now()
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount2.PopCount(uint64(i)))
	}
	time2 = time.Since(start).Microseconds()

	fmt.Printf("popcount v3\n");
	start = time.Now()
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount3.PopCount(uint64(i)))
	}
	time3 = time.Since(start).Microseconds()

	fmt.Printf("popcount v4\n");
	start = time.Now()
	for i := 0; i <= 64; i++ {
		fmt.Printf("%d = %d\n", i, popcount4.PopCount(uint64(i)))
	}
	time4 = time.Since(start).Microseconds()

	fmt.Printf("popcount v1\n");
	fmt.Printf("%d us elapsed\n", time1)
	fmt.Printf("popcount v2\n");
	fmt.Printf("%d us elapsed\n", time2)
	fmt.Printf("popcount v3\n");
	fmt.Printf("%d us elapsed\n", time3)
	fmt.Printf("popcount v4\n");
	fmt.Printf("%d us elapsed\n", time4)
}