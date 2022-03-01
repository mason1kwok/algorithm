## 动态规划

```go
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    // 初始化dp
    dp := make([][]int, m+1)
    for i, _ := range dp {
        dp[i] = make([]int, n+1)
    }
    // 初始化第一列数据
    for i := 1; i <= m; i++ {
        dp[i][0] = i
    }
    // 初始化第一行数据
    for j := 1; j <= n; j++ {
        dp[0][j] = j
    }

    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            
            if word1[i-1] == word2[j-1] {
                // 相等的情况不增加步数
                dp[i][j] = dp[i-1][j-1]
            } else {
                // 通过插入、删除、替换找到最小步数
                dp[i][j] = min(dp[i][j-1]+1, min(dp[i-1][j]+1, dp[i-1][j-1]+1))
            }
        }
    }
    // 返回结果
    return dp[m][n]
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
```

## 复杂度
- 时间复杂度O(mn), m是word1的长度，n是word2的长度。
- 空间复杂度O(mn)。