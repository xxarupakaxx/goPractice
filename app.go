package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println(s)
	s = append(s, 5, 6)
	fmt.Println(s)
}
