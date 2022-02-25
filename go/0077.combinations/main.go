package main

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	track := []int{}
	backtrack(n, k, 1, track)
	return res
}

// 回溯算法
func backtrack(n, k, start int, track []int) {
	if k == len(track) {
		res = append(res, append([]int{}, track...))
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
