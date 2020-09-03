package main

import (
	"fmt"
	"runtime"
)

func main() {
	go fmt.Println("hello")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %d\n", runtime.Version())
}
