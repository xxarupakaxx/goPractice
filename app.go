package main

import "fmt"

func appName() string {
	const A = "GO Application"
	var Version = "1.0"
	return A + " " + Version
}

func main() {
	fmt.Println(appName())

}
