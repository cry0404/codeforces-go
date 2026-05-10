package main

// https://space.bilibili.com/206214
func isGood1(nums []int) bool {
	n := len(nums) - 1
	cnt := make([]int, n+1)
	for _, x := range nums {
		if x > n ||
			x == n && cnt[x] > 1 || // cnt[x] 加一之前 > 1，加一之后 > 2
			x < n && cnt[x] > 0 {   // cnt[x] 加一之前 > 0，加一之后 > 1
			return false
		}
		cnt[x]++
	}
	return true
}

func isGood(nums []int) bool {
	n := len(nums) - 1
	cntN := 0
	for _, x := range nums {
		x = abs(x)
		if x > n ||
			x == n && cntN > 1 ||
			x < n && nums[x] < 0 {
			return false
		}
		if x == n {
			cntN++
		} else {
			nums[x] = -nums[x] // 标记 x 遇到过
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
