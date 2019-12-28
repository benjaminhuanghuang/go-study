package main

import "fmt"

func main() {
	// Create array with 0 value
	var arr1 [5]int

	// Create array with init value
	arr2 := [3]int{1, 2, 3}   // :=需要初值
	arr3 := [...]int{1, 2, 3} // [] is slice

	// 2d array
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

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
