package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value} // can return a pointer to local variable
}

// Create a method for TreeNode
func (node *TreeNode) Print() {
	fmt.Print(node.Value)
}

// Pass by value
func (node *TreeNode) SetValue(value int) {
	node.Value = value
}
