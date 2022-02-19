func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// 将左子树连接到右子树
	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}