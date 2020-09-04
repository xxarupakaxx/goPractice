package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	case <-ch3:
		fmt.Println("ch3")
	default:
		fmt.Println("NO")
	}
}
