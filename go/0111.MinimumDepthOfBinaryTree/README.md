
## BFS

```go
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    stack := []*TreeNode{}
    level := 0
    stack = append(stack, root)

    for len(stack) > 0 {
        level++
        l := len(stack)
        for i := 0; i < l; i++ {
            node := stack[0]
            stack = stack[1:]
            if node.Left == nil && node.Right == nil {
                return level
            }
            if node.Left != nil {
                stack = append(stack, node.Left)
            }
            if node.Right != nil {
                stack = append(stack, node.Right)
            }
        }   
    }

    return level

}
```

## 复杂度
- 时间复杂度O(N), N为二叉树节点数。
- 空间复杂度O(N)