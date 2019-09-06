/*
Go 语言指针不能运算
*/
package main

import "fmt"

func updateValue() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func updateArrayByPointer(arr *[5]int) {
	arr[0] = 100
}

func updateArray() {
	var arr1 [5]int
	updateArrayByPointer(&arr1)
	fmt.Println(arr1)
}

func main() {
	updateValue()
	updateArray()
}
