package main

import "fmt"

func main() {
	arr3 := [...]int{1, 2, 3} // [] is slice

	// iteration
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// using range
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	// index and value
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	// ignore index
	sum := 0
	for _, v := range arr3 {
		sum += v
	}
}
