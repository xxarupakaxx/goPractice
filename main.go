package main

import "fmt"

type Base struct {
	Id int 
	Owner string
}
type A struct {
	Base 
	Name string
	Area string
}
type B struct {
	Base
	Title string
	Bodies []string
}
a := A{
	Base : Base{17,"Taro"},
	Name :"Taro",
	Area : "Tokyo",
}
func main() {
	var p *int
	fmt.Println(p == nil)
	var i int
	a := &i
	fmt.Printf("%T\n", a)
	pp := &a
	fmt.Printf("%T\n", pp)
}
