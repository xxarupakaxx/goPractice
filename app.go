package main

import "fmt"

func main() {
	fmt.Println("hello")

	s := ""
	for _, v := range []string{"A", "B", "C"} {
		s += v

	}

	fmt.Println(s)

}
