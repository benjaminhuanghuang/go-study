package main

import "fmt"

/*
	Slice is a view of array slice本身没有数据
*/
func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// Create slice [2:6)
	s1 := arr[2:6] // 2,3,4,5
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr) // arr was updated

	updateSlice(arr[:]) // update array
	fmt.Println(arr)
}
