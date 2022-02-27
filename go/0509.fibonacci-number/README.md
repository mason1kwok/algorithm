## 动态规划
```go
func fib(n int) int {
    if n == 0 {
        return 0
    }
    dp_i_1, dp_i_2 := 1, 0
    for i := 2; i <= n; i++ {
        dp_i := dp_i_1 + dp_i_2
        dp_i_2, dp_i_1 = dp_i_1, dp_i
    }
    return dp_i_1
}
```

## 复杂度
- 时间复杂度O(n)
- 空间复杂度O(1)