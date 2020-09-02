package main



func main() {
	import "fmt"

	
	var i = 0
	for {
		fmt.println(i)
		i++
		if i == 10 {
			break
		}
	}
}

/*
func appName() string {
	const A = "GO Application"
	var Version = "1.0"
	return A + " " + Version
}

func main() {
	fmt.Println(appName())

}
*/
