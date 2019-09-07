package main

import "./tree"

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}

	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	// Call the extension method
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
