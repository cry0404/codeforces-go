package main

import "slices"

// https://space.bilibili.com/206214
func concatWithReverse(nums []int) []int {
	rev := slices.Clone(nums)
	slices.Reverse(rev)
	return append(nums, rev...)
}
