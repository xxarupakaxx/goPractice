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
func (p *Point) ToString() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type MyInt int

func (p Point) Set(x, y int) {
	p.X = x
	p.Y = y

}

func (m MyInt) Pluse(i int) int {
	return int(m) + 1
}
func main() {
	f := (*Point).ToString
	f(&Point{7, 11})
}
