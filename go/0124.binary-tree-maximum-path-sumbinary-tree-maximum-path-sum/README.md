## 递归
```go
func maxPathSum(root *TreeNode) int {
    sum := math.MinInt32

    var maxGain func(*TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        leftGain := max(0, maxGain(node.Left))
        rightGain := max(0, maxGain(node.Right))

        pathMaxSum := node.Val + leftGain + rightGain

        sum = max(sum, pathMaxSum)

        return node.Val + max(leftGain, rightGain)
    }
    maxGain(root)
    return sum
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## 复杂度
- 时间复杂度O(N)
- 空间复杂度O(N)