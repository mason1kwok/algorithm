## 动态规划

```go
func lengthOfLIS(nums []int) int {
    // 初始化动态规划表
    dp := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    
    res := 0
    for _, v := range dp {
        res = max(res, v)
    }

    return res+1
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
```

## 复杂度
- 时间复杂度O(n^2)
- 空间复杂度O(n)