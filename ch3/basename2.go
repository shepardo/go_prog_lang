
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("a/b/c.go")) // "c"
	fmt.Println(basename("c.d.go")) // "c.d"
	fmt.Println(basename("abc")) // "abc"
}

func basename(s string) string {
  slash := strings.LastIndex(s, "/") // -1  if "/" not found
  s = s[slash+1:]
  if dot := strings.LastIndex(s, "."); dot >= 0 {
    s = s[:dot]
  }
  return s
}
