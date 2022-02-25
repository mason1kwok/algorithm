package main

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	// 初始化
	res = [][]int{}
	track := []int{}
	// 排序
	sort.Ints(nums)
	// 递归
	backtrack(nums, 0, track)
	return res
}

// 回溯算法
func backtrack(nums []int, start int, track []int) {
	// 添加到结果集
	res = append(res, append([]int{}, track...))

	for i := start; i < len(nums); i++ {
		// 判断是元素是否连续
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		// 递归
		backtrack(nums, i+1, track)
		// 撤销选择
		track = track[:len(track)-1]
	}
}
