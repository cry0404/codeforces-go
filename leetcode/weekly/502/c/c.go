package main

import "math/bits"

// https://space.bilibili.com/206214
func countLocalMaximums(matrix [][]int) (ans int) {
	n, m := len(matrix), len(matrix[0])
	wn, wm := bits.Len(uint(n)), bits.Len(uint(m))
	const MX = 8
	// st[i][j][k1][k2] 表示左上角在 (i, j)，右下角在 (i+(1<<k1)-1, j+(1<<k2)-1) 的子矩阵最大值
	st := make([][][MX][MX]int, n)
	for i := range st {
		st[i] = make([][MX][MX]int, m)
	}
	for i, row := range matrix {
		for j, x := range row {
			st[i][j][0][0] = x
		}
	}
	// 单独计算 k1 = 0
	for k2 := 1; k2 < wm; k2++ {
		for i := range n {
			for j := range m - 1<<k2 + 1 {
				st[i][j][0][k2] = max(st[i][j][0][k2-1], st[i][j+1<<(k2-1)][0][k2-1])
			}
		}
	}
	for k1 := 1; k1 < wn; k1++ {
		for k2 := range wm {
			for i := range n - 1<<k1 + 1 {
				for j := range m - 1<<k2 + 1 {
					st[i][j][k1][k2] = max(st[i][j][k1-1][k2], st[i+1<<(k1-1)][j][k1-1][k2])
				}
			}
		}
	}

	// 返回子矩阵最大值
	// 左闭右开，行号范围 [r1, r2)，列号范围 [c1, c2)
	query := func(r1, c1, r2, c2 int) int {
		r1 = max(r1, 0)
		c1 = max(c1, 0)
		r2 = min(r2, n)
		c2 = min(c2, m)
		k1 := bits.Len8(uint8(r2-r1)) - 1
		k2 := bits.Len8(uint8(c2-c1)) - 1
		// 视作四个子矩阵的并集
		return max(st[r1][c1][k1][k2], st[r2-1<<k1][c1][k1][k2], st[r1][c2-1<<k2][k1][k2], st[r2-1<<k1][c2-1<<k2][k1][k2])
	}

	for i, row := range matrix {
		for j, x := range row {
			if x > 0 && max(query(i-x, j-x+1, i+x+1, j+x), query(i-x+1, j-x, i+x, j+x+1)) <= x {
				ans++
			}
		}
	}
	return
}
