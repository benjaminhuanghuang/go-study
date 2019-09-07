package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[:]
	fmt.Println(s1)

	s2 := s1[:5]
	fmt.Println(s2)

	s3 := arr[2:6]
	fmt.Println(s3)

	i := s3[4] // out of ranage error
	fmt.Println(i)

	s4 := s3[3:5] // get value from array
	fmt.Println(s4)
}
