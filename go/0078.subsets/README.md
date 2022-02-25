## 回溯算法

```go
var res [][]int

func subsets(nums []int) [][]int {
    // 初始化
    res = [][]int{}
    track := []int{}
    // 递归
    backtrack(nums, 0, track)
    return res
}

// 回溯算法
func backtrack(nums []int, start int, track []int) {
    // 添加到结果集
    res = append(res, append([]int{}, track...))

    for i := start; i < len(nums); i++ {
        // 做选择
        track = append(track, nums[i])
        // 递归
        backtrack(nums, i+1, track)
        // 撤销选择
        track = track[:len(track)-1]
    }
}
```

## 复杂度
- 时间复杂度O(n^2), 其中n是数组的长度。
- 空间复杂度O(n)