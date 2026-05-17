package main

// https://space.bilibili.com/206214
func isAdjacentDiffAtMostTwo(s string) bool {
	for i := range len(s) - 1 {
		if abs(int(s[i])-int(s[i+1])) > 2 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
