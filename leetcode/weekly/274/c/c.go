package main

import (
	"math"
	"math/bits"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func asteroidsDestroyed1(mass int, asteroids []int) bool {
	slices.Sort(asteroids)
	for _, x := range asteroids {
		if mass < x { // 无法摧毁小行星 x
			return false
		}
		mass += x // 获得这颗小行星的质量
	}
	return true
}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	maxWidth := bits.Len(uint(slices.Max(asteroids)))
	sum := make([]int, maxWidth)
	mn := make([]int, maxWidth)
	for i := range mn {
		mn[i] = math.MaxInt
	}

	for _, x := range asteroids {
		i := bits.Len(uint(x)) - 1
		sum[i] += x
		mn[i] = min(mn[i], x)
	}

	for i, m := range mn {
		if m == math.MaxInt {
			continue
		}
		if mass < m { // 无法摧毁这组的任意小行星
			return false
		}
		mass += sum[i] // 获得这组小行星的质量
	}
	return true
}
