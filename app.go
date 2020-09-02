package main

import "fmt"

func runDefer() {
	defer fmt.Println("defer")
	fmt.Println("de")
}

func main() {
	runDefer()

}
