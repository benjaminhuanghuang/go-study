package main

import "fmt"

/*
	Slice is a view of array slice本身没有数据
*/
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// Create slice [2:6)
	s := arr[2:6] // 2,3,4,5
	fmt.Println(s)

	s = arr[:6] // 0 to 5

	s = arr[2:] // 2 to end

	s = arr[:] // all

}
