package popcount4

/*
Exercise 2.5: The expression x&(x1)
clears the right most non-zero bit of x. Write a version
of PopCount that counts bits by using this fact, and assess its performance.
*/


// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1)
		result++
	}
  return result
}