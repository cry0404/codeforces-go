package main

import "math"

// https://space.bilibili.com/206214
func f(n, k int) int {
	if n < 0 {
		return 0
	}
	x := int(math.Pow(float64(n), 1/float64(k)))
	// 可能 x 的正确值是 6，但算出来的 x = int(5.99999...) = 5
	if pow(x+1, k) <= n { // 避免浮点误差，这里用整数计算 pow
		x++
	}
	return x + 1
}

func countKthRoots(l, r, k int) int {
	return f(r, k) - f(l-1, k)
}

// LC50. Pow(x, n)
func pow(x, k int) int {
	res := 1
	for ; k > 0; k /= 2 {
		if k%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}
