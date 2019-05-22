package main

func printArray(arr *[5]int) {
	arr[0] = 100
}

func arrayDemo() {
	var arr1 [5]int

	printArray(&arr1)

}
