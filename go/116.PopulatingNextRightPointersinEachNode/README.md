## 非递归遍历

```go
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }

    queue := []*Node{}

    queue = append(queue, root)

    for len(queue) > 0 {
        l := len(queue)
        for i := 0; i < l; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            // 填充右侧节点，即此时数组的第一个节点
            if i < l - 1 {
                node.Next = queue[0]
            }

            
        }
    }
    return root
}
```

## 复杂度

- 时间复杂度O(N), N为二叉树节点数。
- 空间复杂度O(1)
