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

func updateSlice(s []int) {

}

func sliceAttributes() {

}

func main() {
	arrayToSlice()
}
