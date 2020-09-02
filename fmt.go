package main

import "fmt"

const One = 1

func one() (int, int) {
	const two = 2
	return One, two
}

func main() {
	x, y := one()
	fmt.Printf("x=%d,y=%d", x, y)
}

/*
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
}


func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}
func main() {
	f := later()

	fmt.Println(f("golang"))
	fmt.Println(f("is"))
	fmt.Println(f("awesome"))
}


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
}

func div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}
func main() {
	q, r := div(19, 7)
	fmt.Printf("%d,%d\n", q, r)
}

func callFunction(f func()) {
	f()
}

func main() {
	callFunction(func())
	fmt.Printf("im a function")
}*/
