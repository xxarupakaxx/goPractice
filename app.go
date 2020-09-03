package main

import "fmt"

func main() {
	defer println("112")

	panic("runtime error")
	fmt.Println("hee")
}
