func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果当前节点值大于目标值， 搜索左子树；否者搜索右子树。
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	}

	return root
}