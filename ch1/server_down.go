// Server1 is a minimal "echo" server.
package main

import (
  "fmt"
  "log"
	"net/http"
	"time"
)

func main() {
  http.HandleFunc("/down", handler) // each request calls handler
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Second)
  fmt.Fprintf(w, "URL.Path = %q %v\n", r.URL.Path, time.Now().Unix())
}
