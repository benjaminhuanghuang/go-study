package main

import (
	"fmt"

	"../../tree"
)

func main() {
	root := tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.TreeNode)
	root.Right.Left.SetValue(4)

	root.Traverse()

	// Pass function to traversal
	nodeCount := 0
	root.TraverseFunc(func(node *tree.TreeNode) {
		nodeCount++
	})
	fmt.Println(nodeCount)
}
