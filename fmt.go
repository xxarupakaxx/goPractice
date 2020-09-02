package main

import "fmt"

/*
var n = 100
var c = complex(1.0, 3)
var a =[5]int{1,2,3,4,5}

func plus(x,y int) int {
	return x + y
}

func main() {
	n = n + 1
	fmt.Printf("n=%d\n", n)
	fmt.Printf("c=%v\n", c)
}


func main() {
	fmt.Printf("10進数=%d, %b, %o,%x\n", 17, 17, 17, 17)
}

var n int
var x,y,z int

var {
	x, y int
	name string
}*/

func div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}
func main() {
	q, r := div(19, 7)
	fmt.Printf("%d,%d\n", q, r)
}
