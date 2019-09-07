package main

import "../../tree"

func main() {
	root := tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{Value: 5, Left: nil, Right: nil}
	root.Right.Left = new(tree.TreeNode)
	root.Right.Left.SetValue(4)

	root.Traverse()
}
