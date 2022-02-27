package main

var memo map[int]int

func coinChange(coins []int, amount int) int {
    memo = map[int]int{}
    return dp(coins, amount)
}

func dp(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }
    // 查备忘录
    if v, ok := memo[amount]; ok {
        return v
    }

    res := math.MaxInt

    for _, coin := range coins {
        // 计算子问题
        sub := dp(coins, amount - coin)
        if sub == -1 {
            continue
        }
        // 选择最小的结果
        res = min(res, sub+1)
    }

    if res == math.MaxInt {
        memo[amount] = -1
    } else {
        memo[amount] = res
    }

    return memo[amount]
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}