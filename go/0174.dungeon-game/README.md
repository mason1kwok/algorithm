## 回溯法

```go
var memo [][]int

func calculateMinimumHP(dungeon [][]int) int {
    m := len(dungeon)
    n := len(dungeon[0])
    // 初始化备忘录
    memo = make([][]int, m)
    for i, _ := range memo {
        memo[i] = make([]int, n)
        for j := 0; j < n; j++ {
            memo[i][j] = -1
        }
    }
    // 从0,0开始走
    return dp(dungeon, 0, 0)
}

func dp(grid [][]int, i, j int) int {
    m := len(grid)
    n := len(grid[0])

    // base case
    if i == m-1 && j == n-1 {
        if grid[i][j] >= 0 {
            return 1
        }
        return -grid[i][j] + 1
    }


    if i == m || j == n {
        return math.MaxInt
    }
    
    if memo[i][j] != -1 {
        return memo[i][j]
    }

    // 状态转移逻辑
    res := min(dp(grid, i, j+1), dp(grid, i+1, j)) - grid[i][j]

    memo[i][j] = res
    if res <= 0 {
        memo[i][j] = 1
    }
    // 返回结果
    return memo[i][j]
}

func min(a, b int) int {
    if b > a {
        return a
    }
    return b
}
```

## 复杂度

- 时间复杂度O(mn)
- 空间复杂度O(mn)