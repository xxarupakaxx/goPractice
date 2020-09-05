package main

import "fmt"

type (
	IntPair     [2]int
	Strings     []string
	AreaMap     map[string][2]float64
	IntsChannel chan []int
	CallBack    func(i int) int
)

func Sum(ints []int, CallBack CallBack) int {
	var sum int
	for _, i := range ints {
		sum += i

	}
	return CallBack(sum)
}
func main() {
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n)
}
