package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func countVowelSubstrings(word string) (ans int) {
	cnt := map[byte]int{}
	start, left := 0, 0
	for i, ch := range word {
		if strings.IndexRune("aeiou", ch) < 0 {
			clear(cnt) // 重置
			start, left = i+1, i+1
			continue
		}

		// 做法类似 1358. 包含所有三种字符的子字符串数目（哈希表写法）
		cnt[byte(ch)]++
		for len(cnt) == 5 { // 窗口包含所有元音
			out := word[left]
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
			left++
		}
		ans += left - start
	}
	return
}
