package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 0)
	fmt.Printf("len = %d , cap = %d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len = %d , cap = %d\n", len(s), cap(s))
	s = append(s, []int{2, 3, 4}...)
	fmt.Printf("len = %d , cap = %d\n", len(s), cap(s))
}
