package main

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}

	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res
}