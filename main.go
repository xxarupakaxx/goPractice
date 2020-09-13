package main

import (
	"fmt"
	"math"
)

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}

type Point struct{ X, Y int }

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}
func main() {
	p := Point{1, 2}
	fmt.Println(p.Distance(&Point{0, 0}))

}
