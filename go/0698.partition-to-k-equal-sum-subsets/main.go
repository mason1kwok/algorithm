var memo map[int]bool

func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	memo = map[int]bool{}
	target := sum / k
	return backtrack(nums, 0, k, 0, 0, target)
}

// 递归
func backtrack(nums []int, bucket, k, start, used, target int) bool {
	if k == 0 {
		return true
	}

	if bucket == target {
		res := backtrack(nums, 0, k-1, 0, used, target)
		memo[used] = res
		return res
	}

	if res, ok := memo[used]; ok {
		return res
	}

	for i := start; i < len(nums); i++ {
		if ((used >> i) & 1) == 1 {
			continue
		}
		if nums[i]+bucket > target {
			continue
		}

		used |= 1 << i
		bucket += nums[i]
		if backtrack(nums, bucket, k, i+1, used, target) {
			return true
		}
		bucket -= nums[i]
		used ^= 1 << i

	}
	return false
}