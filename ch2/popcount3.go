package popcount3

/*
Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
bit positions, testing the rightmost bit each time. Compare its performance to the table 
lookup version.
*/

// pc[i] is the population count of i.
var pc [256]byte

func init() {
  for i := range pc {
    pc[i] = pc[i/2] + byte(i&1)
  }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
  result := 0
  for i := 0; i <= 63; i++ {
    if x & 1 == 1 {
			result++
		}
		x = x >> 1
  }
  return result
    /*
  return int(pc[byte(x>>(0*8))] +
    pc[byte(x>>(1*8))] +
    pc[byte(x>>(2*8))] +
    pc[byte(x>>(3*8))] +
    pc[byte(x>>(4*8))] +
    pc[byte(x>>(5*8))] +
    pc[byte(x>>(6*8))] +
    pc[byte(x>>(7*8))])*/
}