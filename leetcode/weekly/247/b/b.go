package main

// github.com/EndlessCheng/codeforces-go
func rotateGrid(grid [][]int, k int) [][]int {
	m0, n0 := len(grid), len(grid[0])

	// 从外到内枚举圈
	for i := range min(m0, n0) / 2 {
		m, n := m0-i*2-1, n0-i*2-1 // 这一圈的行数-1、列数-1

		// 返回这一圈顺时针下标 p 对应 grid 的位置 (x, y)
		index := func(p int) (x, y int) {
			// 左上角在 (i, i)
			if p < n {
				return i, i + p
			}
			if p < n+m {
				return i + p - n, i + n
			}
			if p < n*2+m {
				return i + m, i - p + n*2 + m
			}
			return i - p + (n+m)*2, i
		}

		reverse := func(l, r int) {
			for l < r {
				x1, y1 := index(l)
				x2, y2 := index(r)
				grid[x1][y1], grid[x2][y2] = grid[x2][y2], grid[x1][y1]
				l++
				r--
			}
		}

		// 189. 轮转数组（改成向左轮转）
		size := (m + n) * 2
		shift := k % size
		reverse(0, shift-1)
		reverse(shift, size-1)
		reverse(0, size-1)
	}

	return grid
}
