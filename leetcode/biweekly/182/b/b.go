package main

import "strings"

// https://space.bilibili.com/206214
func minFlips(s string) int {
	n := len(s)
	c0 := strings.Count(s, "0")
	c1 := n - c0 - 1
	if s[0] == '1' && s[n-1] == '1' {
		c1--
	}
	return min(c0, max(c1, 0))
}
