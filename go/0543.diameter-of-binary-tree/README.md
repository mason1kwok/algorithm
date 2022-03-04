## 递归

```go
func diameterOfBinaryTree(root *TreeNode) int {
    ans := 1
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        ans = max(ans, left+right+1)
        return max(left, right) + 1
    }
    dfs(root)
    return ans - 1
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## 复杂度
- 时间复杂度O(n)
- 空间复杂度O(1)
