package main

import (
	"fmt"
	"math"
)
type MyError struct{
	Message string
	ErrCode int
}
func (e *MyError) Error() string {
	return e.Message
}
func RaiseError()error{
	return MyError{Message:"error",ErrCode:1234}
}
err
type User struct {
	Id   int    "ユーザー"
	Name string "名前"
	Age  uint   "年齢"
}

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

}
