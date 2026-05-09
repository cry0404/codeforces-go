package main

import (
	"slices"
	"strconv"
)

// https://space.bilibili.com/206214
func separateDigits1(nums []int) (ans []int) {
	for _, x := range nums {
		for _, ch := range strconv.Itoa(x) {
			ans = append(ans, int(ch-'0'))
		}
	}
	return
}

func separateDigits(nums []int) (ans []int) {
	for _, x := range slices.Backward(nums) {
		for ; x > 0; x /= 10 {
			ans = append(ans, x%10)
		}
	}
	slices.Reverse(ans)
	return
}
