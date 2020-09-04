package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	for i := range ch {
		fmt.Println(i)
	}
}
