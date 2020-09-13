package main

import "fmt"

func plus(x, y int) int {
	return x + y
}
func returnFunc() func() {
	return func() {
		fmt.Println("I,n")
	}
}

func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s

	}
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Println(v)
			case string:
				fmt.Println(v)
			default:
				fmt.Println(v)
			}
		}
	}()
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)
	go receiver(ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	/*r := '松'
	fmt.Printf("%v", r)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a)

	fmt.Printf("%v\n", func(x, y int) int { return x + y })
	fmt.Printf("%v\n", func(x, y int) int { return x + y }(2, 4))
	f := returnFunc()
	f()*/

	/*panic("runtime")

	defer func() {
		if x := recover(); x !=nil {
			fmt.Println(x)
		}
	}()
	paic("aa")*/

	select{
	case e1 := <-ch1
	case e2 := <-ch2
	default:
	}

}
