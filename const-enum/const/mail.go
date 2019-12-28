/*
const variable 不要大写，大写表明是变量在package外可以访问

*/
package main

import (
	"fmt"
	"math"
)

func consts() {
	const filename = "abc.txt"
	/*********************************************
	// 如果指定类型会导致 error: cannot use a * a + b * b (type int) as type float64 in argument to math.Sqrt
	// const a, b int = 3, 4
	***********************************************/
	// 如果不指定类型，常量的类型不确定，在使用时只会做文本替换，不管类型
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(filename, c)
}

func main() {
	consts()
}
