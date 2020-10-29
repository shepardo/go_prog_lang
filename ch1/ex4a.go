
/*
Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
*/

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	counts := make(map[string]int)
	filesXref := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesXref, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//fmt.Fprintf(os.Stdout, "File: %s\n", arg)
			countLines(f, counts, filesXref, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			//fmt.Printf("in files %s\n", strings.Join(filesXref[line], ",") )
			fmt.Printf("in files %s\n", listToString(filesXref[line]))
		}
	}
}

func listToString(m map[string]int) string {
	s := ""
	sep := ""
	for k, _ := range m {
		s += sep + k
		sep = ","
	}
	return s
}

func countLines(f *os.File, counts map[string]int, files map[string]map[string]int, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		word:= input.Text()
		counts[word]++
		if _, ok := files[word]; !ok {
			files[word] = make(map[string]int)
		}
		if _, ok := files[word][filename]; !ok {
			files[word][filename] = 1
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
