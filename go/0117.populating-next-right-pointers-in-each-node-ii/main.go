package main

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{root}

	for len(q) > 0 {
		n := len(q)
		var pre *Node
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]

			if pre != nil {
				pre.Next = cur
			}

			pre = cur

			if cur.Left != nil {
				q = append(q, cur.Left)
			}

			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}

	return root
}
