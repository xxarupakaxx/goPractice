package main

import "fmt"

func main() {
	m := map[string]int{
		"Apple":  88,
		"Banana": 107,
		"vherry": 46,
	}
	m["Grape"] = 66
	m["Lemon"] = 16
	m["Orange"] = 44
	m["Oineapple"] = 73
	for k, v := range m {
		fmt.Printf("%v => %v\n", k, v)
	}
}
