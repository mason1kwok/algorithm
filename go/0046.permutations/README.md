## 回溯算法

```go
var res [][]int

func permute(nums []int) [][]int {
    res = [][]int{}
    track := []int{}
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
        if listContains(track, nums[i]) {
            continue
        }
        // 做选择
        track = append(track, nums[i])
        // 递归
        backtrack(nums, track)
        // 撤销选择
        track = track[:len(track)-1]
    }
}

func listContains(a []int, b int) bool {
    for _, v := range a {
        if v == b {
            return true
        }
    }
    return false
}
```

## 复杂度

- 时间复杂度O(n * n!),n是数组的长度。
- 空间复杂度O(n)