
## 递归

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return targetSum == root.Val
    }

    return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}
```

## 复杂度
- 时间复杂度O(n)
- 空间复杂度O(1)