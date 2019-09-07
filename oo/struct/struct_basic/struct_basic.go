package main

import "fmt"

type point struct {
	i, j int
}

var dir = []point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

type Rect struct {
	width, length float64
}

func doubleArea(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func rectDemo() {
	// Create rect
	var rect = Rect{width: 100, length: 200}

	// Create rect
	/*
		如果知道结构体成员定义的顺序，也可以不使用key:value的方式赋值，直接按照结构体成员定义的顺序给它们赋值。
	*/
	rect = Rect{100, 200}
	fmt.Println("Width:", rect.width, "* Length:", rect.length, "= Area:", rect.width*rect.length)

	// Pass struct as value(copy)
	fmt.Println(doubleArea(rect))
	fmt.Println("Width:", rect.width, "Length:", rect.length)
}

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func main() {
	var root TreeNode
	root = TreeNode{value: 3}
	root.left = &TreeNode{}
	root.right = &TreeNode{5, nil, nil}
	root.right.left = new(TreeNode)

	nodes := []TreeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
