package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1805F2(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)

	m := min(n, 100)
	ans := 0
	for i := range n {
		ans = (ans*2 + a[0]) % mod

		for j := 1; j < m; j++ {
			a[j] -= a[0]
		}
		a[0] = 0

		b := []int{}
		for k := range m {
			for j := k + 1; (k+1)*j <= m*2 && j < m; j++ {
				b = append(b, a[k]+a[j])
			}
		}
		slices.Sort(b)

		m = min(m, n-i-1)
		copy(a[:m], b[:m])
	}
	Fprint(out, ans)
}

//func main() { cf1805F2(bufio.NewReader(os.Stdin), os.Stdout) }
