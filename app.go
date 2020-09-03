package main

import "fmt"

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("panic")

	fmt.Println("jhe")
}
