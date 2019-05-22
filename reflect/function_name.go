package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf(opName)

	return op(a, b)
}

func main() {
	fmt.Println(apply(func(a int, b int) int {
		return a + b
	}, 3, 4))
}
