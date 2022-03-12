## DFS实现
```go
var result bool

func isBalanced(root *TreeNode) bool {
    result = true
    dfs(root)
    return result
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := dfs(root.Left)
    right := dfs(root.Right)
    if abs(left - right) > 1 {
        result = false
    }
    return max(left, right) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
```

## 复杂度
- 时间复杂度O(N)
- 空间复杂度O(1)