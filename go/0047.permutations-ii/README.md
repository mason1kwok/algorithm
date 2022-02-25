## 回溯

```go
var (
    res [][]int
    visited map[int]bool
)

func permuteUnique(nums []int) [][]int {
    res = [][]int{}
    track := []int{}
    visited = map[int]bool{}
    sort.Ints(nums)
    backtrack(nums, track)
    return res
}

// 回溯算法
func backtrack(nums []int, track []int) {
    // 符合条件，加入结果集
    if len(nums) == len(track) {
        res = append(res, append([]int{}, track...))
        return
    }

    // 循环递归调用
    for i := 0; i < len(nums); i++ {
        // 现在的数组是否含有当前数
        if visited[i] || (i > 0 && nums[i] == nums[i-1] && visited[i-1]) {
            continue
        }

        // 做选择
        track = append(track, nums[i])
        visited[i] = true
        // 递归
        backtrack(nums, track)
        // 撤销选择
        visited[i] = false
        track = track[:len(track)-1]
    }
}
```

## 复杂度
- 时间复杂度O(n * n!), n为数组长度
- 空间复杂度O(n)