package main

import (
	"fmt"

	"./animals"
)

func main() {

	fmt.Println(appName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
