package main

import "fmt"

func demo() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func main() {
	demo()
}
