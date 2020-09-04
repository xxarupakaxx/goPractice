package main

import "fmt"

func main() {
	var i int
	p := &i
	fmt.Printf("%T\n", p)
	pp := &p
	fmt.Printf("%T\n", pp)
}
