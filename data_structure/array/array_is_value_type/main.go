package main

import "fmt"

/*
	array is value type, should define the length of the array in parameter
	[10]int is differnt type with [11]int

	function all will copy the array
*/
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayPointer(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//arr3 := [...]int{1, 2, 3}
	//printArray(arr3) // error!
	arr := [...]int{1, 2, 3, 4, 5}

	printArray(arr)
	fmt.Println(arr)

	printArrayPointer(&arr)
	fmt.Println(arr)
}
