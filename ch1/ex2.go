/*
Exercise 1.2: Modif y the echo prog ram to print the index and value of each of its arguments,
one per line.
*/
package main
import (
  "fmt"
	"os"
	"strconv"
)
func main() {
  s := "" //, sep := "", ""
  for idx, arg := range os.Args[1:] {
    s += strconv.Itoa(idx) + " " + arg + "\n";
  }
  fmt.Println(s)
}
