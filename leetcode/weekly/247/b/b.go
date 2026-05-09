package main

// github.com/EndlessCheng/codeforces-go
var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上

func rotateGrid(grid [][]int, k int) [][]int {
	m0, n0 := len(grid), len(grid[0])
	a := make([]int, 0, (m0+n0-2)*2) // 预分配空间

	// 从外到内枚举圈
	for i := range min(m0, n0) / 2 {
		m, n := m0-i*2, n0-i*2 // 这一圈的行数和列数
		x, y := i, i // 这一圈的左上角
		a = a[:0]
		for _, dir := range dirs {
			for range n - 1 {
				a = append(a, grid[x][y])
				x += dir[0]
				y += dir[1]
			}
			m, n = n, m // 见 54. 螺旋矩阵 我的题解
		}

		shift := k % len(a)
		a = append(a[shift:], a[:shift]...)

		// 注意此时 (x, y) 又回到了左上角
		j := 0
		for _, dir := range dirs {
			for range n - 1 {
				grid[x][y] = a[j]
				j++
				x += dir[0]
				y += dir[1]
			}
			m, n = n, m
		}
	}

	return grid
}
