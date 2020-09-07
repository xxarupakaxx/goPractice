package main

import "fmt"

type MyStrict struct {
	a    string
	b, c int
}

func main() {
	var st MyStrict
	st.a = "hoge"
	st.b = 1
	st.c = 0
	fmt.Println(st.a, st.b, st.c)
}
