package main

import "strings"

// https://space.bilibili.com/206214
func countWordOccurrences(chunks []string, queries []string) []int {
	s := strings.Join(chunks, "")
	n := len(s)
	cnt := map[string]int{}

	for i := 0; i < n; i++ {
		if s[i] == ' ' || s[i] == '-' {
			continue
		}
		start := i
		// 遇到 ' ' 或者 "--" 或者 "- " 时，跳出循环
		for i < n && s[i] != ' ' && (s[i] != '-' || i < n-1 && s[i+1] != '-' && s[i+1] != ' ') {
			i++
		}
		cnt[s[start:i]]++
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = cnt[q]
	}
	return ans
}
