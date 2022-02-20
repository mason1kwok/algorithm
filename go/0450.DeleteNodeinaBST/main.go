package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 获取右子树的最小节点
		minNode := minNode(root.Right)
		// 删除右子树最小节点
		root.Right = deleteNode(root.Right, minNode.Val)
		// 用最小节点替换root节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func minNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
