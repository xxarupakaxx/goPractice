package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	fruits := [3]string{"apple", "baana", "cherry"}

	for i, s := range fruits {
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	anything(1)
	anything(3.14)
	anything(1 + 3.0i)
	anything("海")

}
