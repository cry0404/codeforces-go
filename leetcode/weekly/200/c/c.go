package main

// github.com/EndlessCheng/codeforces-go
func minSwaps(grid [][]int) (ans int) {
	// 预处理每一行的尾零个数
	n := len(grid)
	tailZeros := make([]int, n)
	for i, row := range grid {
		tailZeros[i] = n
		for j := n - 1; j >= 0; j-- {
			if row[j] == 1 {
				tailZeros[i] = n - 1 - j
				break
			}
		}
	}

next:
	for i := range n - 1 {
		needZeros := n - 1 - i
		for j := i; j < n; j++ {
			if tailZeros[j] >= needZeros {
				ans += j - i
				// 从 j 换到 i，原来 [i, j-1] 中的数据全体右移一位
				copy(tailZeros[i+1:j+1], tailZeros[i:j])
				continue next
			}
		}
		return -1
	}
	return
}
