package main

import "fmt"

type Rect struct {
	width, length float64
}

func main() {
	// Create rect
	var rect = Rect{width: 100, length: 200}

	// Create rect
	/*
		然如果你知道结构体成员定义的顺序，也可以不使用key:value的方式赋值，直接按照结构体成员定义的顺序给它们赋值。
	*/
	rect = Rect{100, 200}
	fmt.Println("Width:", rect.width, "* Length:", rect.length, "= Area:", rect.width*rect.length)

	// Pass struct as value(copy)
	fmt.Println(doubleArea(rect))
	fmt.Println("Width:", rect.width, "Length:", rect.length)
}

func doubleArea(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}
