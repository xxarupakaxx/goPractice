package main

import (
	"fmt"

	"./animals"
)

func main() {

	fmt.Println(App())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
