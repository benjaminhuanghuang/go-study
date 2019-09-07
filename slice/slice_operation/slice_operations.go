package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("s= %v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}

func main() {
	s1 := []int{1, 2, 3, 4}
	printSlice(s1)
	// create slice with length and capacity
	s2 := make([]int, 16, 30)
	printSlice(s2)

	// copy
	copy(s2, s1)
	printSlice(s2)

	// Delete
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// pop from front
	s2 = s2[1:]

	// pop from back
	s2 = s2[:len(s2)-1]
}
