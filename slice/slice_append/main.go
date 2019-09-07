package main

import "fmt"

/*
	append slice，底层array会被改变
	append如果超过slice的cap，系统会重新分配更大的底层array
*/
func printSlice(s []int) {
	fmt.Printf("s= %v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}

	// Create slice [2:6)
	s1 := arr[2:6] // 2,3,4,5
	printSlice(s1)
	s1 = append(s1, 10)
	printSlice(s1)
	fmt.Println(arr) // array was changed

	s1 = append(s1, 11)
	printSlice(s1)
	fmt.Println(arr) // arr keep values, system create a new array for slice
}
