/*
Exercise 1.12: Modify the Lissajous ser ver to read parameter values from the URL. For example,
you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the
number of cycles to 20 instead of the defau lt 5. Use the strconv.Atoi function to convert the
string parameter into an integer. You can see its documentation with go doc strconv.Atoi.
*/

// Server3 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"image/gif"
	"image/draw"
	"io"
	"math"
	"math/rand"
//	"os"
	"strconv"
)

var mu sync.Mutex
var count int

//var palette = []color.Color{color.White, color.Black}

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},		// Green
	color.RGBA{0xff, 0x00, 0x00, 0xff},		// Red
	color.RGBA{0x00, 0x00, 0xff, 0xff},		// Blue
	color.RGBA{0xff, 0xff, 0x00, 0xff},		// Yellow 
	color.RGBA{0xff, 0x00, 0xff, 0xff}, 	// Magenta
	color.RGBA{0x00, 0xff, 0xff, 0xff},		// Cyan
}
const (
//	whiteIndex = 0 // first color in palette
	blackIndex = 0 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	/*
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	*/
	bgIndexStr := r.URL.Query().Get("bgColor")
	if bgIndexStr != "" {
		bgIndex, err := strconv.Atoi(bgIndexStr)
		if err == nil {
			lissajous(w, bgIndex)
		}
	} else {
		lissajous(w, 1)
	}
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mu.Unlock()
}

func lissajous(out io.Writer, idxBg int) {
	const (
		cycles =5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay =8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	//idxBg := 1	// Background color
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		bgColor := palette[idxBg]
		var _ = bgColor
		var _ = draw.Draw
		draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)
		idxBg = ((idxBg + 1) % len(palette) - 1) + 1
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}





