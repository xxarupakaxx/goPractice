package main

import "fmt"

type Point struct{ x, y int }

func (p *point) Render() {
	fmt.Printf("<%v,%v>\n", p.X, p.Y)
}
func main() {
	p := &Point{X: 5, Y: 12}

	p.Render()
}
