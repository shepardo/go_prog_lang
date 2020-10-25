/*
Exercise 1.3: Experiment to measure the dif ference in running time bet ween our potential ly
inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates par t of the
time package, and Sec tion 11.4 shows how to write benchmark tests for systematic performance
evaluation.)
*/
package main

import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf("%dms elapsed\n", time.Since(start).Milliseconds())
}