package main

import "fmt"

// using range with slice
func findMax() {
	nums := []int{3, 4, 5, 7, 8, 8}

	max := nums[0]

	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	fmt.Println(max)
}

func arrayToSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 左闭右开
	s := arr[2:6] // 2,3,4,5
	fmt.Println(s)

	s = arr[:6]
	fmt.Println(s)

	s = arr[2:]
	fmt.Println(s)

	s = arr[:]
	fmt.Println(s)

	// s = arr[9]    // Error
}

// update slice 会修改slice背后的array
func updateSlice(s []int) {
	s[0] = 100
}

func updateSliceDemo() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:]
	s2 := arr[:]
	updateSlice(s1)
	fmt.Println("s1 = ", s1)   // s1 updated
	fmt.Println("s2 = ", s2)   // s2 updated
	fmt.Println("arr = ", arr) // arr updated
}

func reslice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:]
	s2 := s1[:4]
	fmt.Println("s1 = ", s1) // s1 updated
	fmt.Println("s2 = ", s2) // s2 updated
}

func sliceExtend() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//s1 := arr[2:100] // Error, out of bounds
	s1 := arr[2:6]
	s2 := s1[3:6] // extending, 6 is the bound of array
	//s2 = s1[3:50]   // Error, out of range
	fmt.Println("s1 = ", s1) // s1 updated
	fmt.Println("s2 = ", s2) // s2 updated
}

func sliceAttributes() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[3:6]

	fmt.Printf("s1 = %v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func addElementsToSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[:]
	s3 := append(s2, 10)
	s4 := append(s3, 11) // view to a new array
	s5 := append(s4, 12) // view to a new array
	fmt.Println(s3, s4, s5)
	fmt.Println(arr)
}

func main() {
	//updateSliceDemo()

	//reslice()

	//sliceExtend()
	addElementsToSlice()
}
