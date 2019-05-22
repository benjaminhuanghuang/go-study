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

func sliceAttributes() {

}

func main() {

}
