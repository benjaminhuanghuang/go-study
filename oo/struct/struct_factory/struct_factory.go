package main

import "fmt"

// struct 没有构造函数，
type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(val int) *treeNode {
	// sfactory method返回局部变量的地址，在外部依然有效。Go语言来决定在heap还是在
	return &treeNode{value: val}
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
