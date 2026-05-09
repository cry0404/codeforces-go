package main

import "math"

// github.com/EndlessCheng/codeforces-go
func minMoves(nums []int, limit int) int {
	n := len(nums)
	diff := make([]int, limit*2+2)
	for i, x := range nums[:n/2] {
		y := nums[n-1-i]
		l := min(x, y) + 1
		r := max(x, y) + limit
		diff[l]--
		diff[x+y]--
		diff[x+y+1]++
		diff[r+1]++
	}

	ans := math.MaxInt
	sum := n
	for _, d := range diff[2 : limit*2+1] {
		sum += d
		ans = min(ans, sum)
	}
	return ans
}
