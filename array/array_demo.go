package main

import "fmt"

func updateArray(arr [3]int) {
	fmt.Println("--updateArray--")
	arr[0] = 100
}

func updateArrayByPointer(arrPointer *[3]int) {
	fmt.Println("--updateArrayByPointer--")
	arrPointer[0] = 100
}

func updateArrayBySlice(slice []int) {
	fmt.Println("--updateArrayBySlice--")
	slice[0] = 100
}

func main() {
	var arr1 [5]int

	arr2 := [3]int{1, 2, 3} // :=需要初值
	arr3 := [...]int{1, 2, 3}

	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	// Loop 1
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// updateArray(arr1)  // Error!  [3]int is different with [5]int

	// array is copied and passed by Value !  no changes to arr3
	updateArray(arr3)
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	updateArrayByPointer(&arr3)
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	updateArrayBySlice(arr3[:])
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
