package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	flag := true
	res := [][]int{}

	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if flag {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, append([]int{}, level...))
		flag = !flag
	}

	return res
}
