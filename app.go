package main

type Person struct {
	Id   int
	Name string
	Area string
}

func main() {
	p := new(Person)

	p.Id
	p.Name
	p.Area
}
