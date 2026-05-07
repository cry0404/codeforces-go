package main

// github.com/EndlessCheng/codeforces-go
func canReach(s string, minJump, maxJump int) bool {
	n := len(s)
	sum := make([]int, n+1) // f 的前缀和
	sum[1] = 1 // f[0] = true
	for j := 1; j < n; j++ {
		sum[j+1] = sum[j]
		if j >= minJump && s[j] == '0' && sum[j-minJump+1] > sum[max(j-maxJump, 0)] {
			sum[j+1]++ // f[j] = true
		}
	}
	return sum[n] > sum[n-1] // f[n-1] == true
}
