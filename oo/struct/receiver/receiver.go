package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// define method for treeNode, node is passed by value
// root.print()
func (node treeNode) print() {
	fmt.Println(node.value)
}

// pass by value
func (node treeNode) setValue(val int) {
	node.value = val
}

// pass pointer
func (node *treeNode) setValueByPointer(val int) {
	node.value = val
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)

	root.print()

	root.right.left.setValue(100) // pass by value
	root.right.left.print()       // doest not change

	root.right.left.setValueByPointer(100) // pass pointer
	root.right.left.print()                // value changed

	pRoot := &root
	pRoot.setValue(200) // pass value
	pRoot.print()       // doest not change

	pRoot.setValueByPointer(200) // pass pointer
	pRoot.print()                // value changed
}
