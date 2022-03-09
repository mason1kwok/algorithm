## 递归
```go
func sortedArrayToBST(nums []int) *TreeNode {
    return dfs(nums, 0, len(nums) - 1)
}

func dfs(nums []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := (left + right + 1) / 2
    root := &TreeNode{Val: nums[mid]}
    root.Left = dfs(nums, left, mid-1)
    root.Right = dfs(nums, mid+1, right)
    return root
}
```

## 复杂度

- 时间复杂度O(n)
- 空间复杂度O(1)