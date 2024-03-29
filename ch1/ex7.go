
/*
Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst. Use it
instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
*/

// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
//	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		bytes, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
		fmt.Printf("%d bytes written.\n", bytes)
	}
}