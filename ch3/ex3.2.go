/*
Exercise 3.2: Experiment with visualizations of other functions from the math package. Can
you produce an egg box, moguls, or a saddle?

For more ideas, see https://www.google.com/imgres?imgurl=https%3A%2F%2Fmathinsight.org%2Fmedia%2Fapplet%2Fimage%2Flarge%2Fsurface_sinxy.png&imgrefurl=https%3A%2F%2Fmathinsight.org%2Fsurface_graph_function&tbnid=Emy3wgF9t9sSCM&vet=12ahUKEwjW75uo4fzuAhUJ_6wKHd3BCQsQMygFegQIARA4..i&docid=v2_jJwxGsljjVM&w=425&h=316&q=what%20is%20a%20mogul%20in%20trigonometric%20graphs&client=firefox-b-d&ved=2ahUKEwjW75uo4fzuAhUJ_6wKHd3BCQsQMygFegQIARA4
*/

// Surface computes an SVG rendering of a 3D surface function.
package main
import (
	"fmt"
	"math"
	"os"
	//"log"
)

const (
	width, height 		= 600, 320 // canvas size in pixels
	cells 						= 100 // number of grid cells
	xyrange 					= 30.0 // axis ranges (-xyrange..+xyrange)
	xyscale 					= width / 2 / xyrange // pixels per x or y unit
	zscale 						= height * 0.1 // pixels per z unit
	angle 						= math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; strokewidth:0.7' " +
		"width='%d' height='%d'>", width, height)
	//l := log.New(os.Stderr, "", 0)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			if math.IsNaN(ay) {
				fmt.Fprintf(os.Stderr, "ay is NaN for %d, %d\n", i + 1, j)
				continue
			}
			bx, by := corner(i, j)
			if math.IsNaN(by) {
				fmt.Fprintf(os.Stderr, "by is NaN for %d, %d\n", i, j)
				continue
			}
			cx, cy := corner(i, j+1)
			if math.IsNaN(cy) {
				fmt.Fprintf(os.Stderr, "cy is NaN for %d, %d\n", i, j+1)
				continue
			}
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(dy) {
				fmt.Fprintf(os.Stderr, "dy is NaN for %d, %d\n", i+1, j+1)
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy).
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func mogul(x, y float64) float64 {
	return math.Abs(math.Sin(x) * math.Sin(y))
}

func f(x, y float64) float64 {
	//r := math.Hypot(x, y) // distance from (0,0)
	//return math.Log2(r) / r
	//return math.Sin(x) * math.Sin(y)
	//return eggbox(x, y)
	//return saddle(x, y)
	return mogul(x, y)
}