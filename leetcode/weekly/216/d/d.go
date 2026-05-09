package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func minimumEffort1(tasks [][]int) (ans int) {
	slices.SortFunc(tasks, func(a, b []int) int {
		return (b[1] - b[0]) - (a[1] - a[0]) // 按照 minimum - actual 从大到小排序
	})

	s := 0 // 累计耗费的能量
	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		// 题目要求 E0 - s >= minimum，即 E0 >= s + minimum
		// 由此可以得到 n 个关于 E0 的下界，所有下界的最大值即为答案
		ans = max(ans, s+minimum)
		s += actual
	}
	return
}

func minimumEffort(tasks [][]int) (e int) {
	slices.SortFunc(tasks, func(a, b []int) int {
		return (a[1] - a[0]) - (b[1] - b[0]) // 按照 minimum - actual 从小到大排序
	})

	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		// 完成 t 之后的能量为 e，那么完成 t 之前的能量为 e+actual，同时该能量必须至少为 minimum
		e = max(e+actual, minimum)
	}
	return
}
