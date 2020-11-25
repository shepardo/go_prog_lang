// Echo4 prints its commandline arguments.

/*
$ ./echo4 a bc def
a bc def
$ ./echo4 s
/ a bc def
a/bc/def
$ ./echo4 n
a bc def
a bc def$
$ ./echo4 help
Usage of ./echo4:
*/
package main
import (
  "flag"
  "fmt"
  "strings"
)
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }
}