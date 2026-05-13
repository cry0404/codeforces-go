package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1285F(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }

	var n, ans, cnt int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	slices.Sort(a)
	a = slices.Compact(a)
	n = len(a)

	for i := n - 1; i >= 0; i-- {
		if i >= n-30 || cnt < 30 && (a[i]^a[n-1])&1 > 0 {
			for j := n - 1; j >= 0; j-- {
				ans = max(ans, lcm(a[i], a[j]))
			}
			if i < n-30 {
				cnt++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1285F(bufio.NewReader(os.Stdin), os.Stdout) }
