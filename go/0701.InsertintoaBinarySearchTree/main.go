package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到新节点的位置，插入
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
