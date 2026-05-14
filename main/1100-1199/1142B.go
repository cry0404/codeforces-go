package main

import (
	"bytes"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1142B(in io.Reader, out io.Writer) {
	var n, m, q, v, l, r int
	Fscan(in, &n, &m, &q)
	p := make([]int, n)
	pre := make([]int, n+1)
	for i := range p {
		Fscan(in, &p[i])
		if i > 0 {
			pre[p[i]] = p[i-1]
		}
	}
	pre[p[0]] = p[n-1]

	jumpL := make([][18]int, m+1)
	last := make([]int, n+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &v)
		jumpL[i][0] = last[pre[v]]
		last[v] = i
	}

	for j := range 17 {
		for i := range jumpL {
			jumpL[i][j+1] = jumpL[jumpL[i][j]][j]
		}
	}

	maxL := make([]int, m+1)
	for i := 1; i <= m; i++ {
		cur := i
		for s := uint(n - 1); s > 0; s &= s - 1 {
			cur = jumpL[cur][bits.TrailingZeros(s)]
		}
		maxL[i] = max(maxL[i-1], cur)
	}

	ans := bytes.Repeat([]byte{'0'}, q)
	for i := range q {
		Fscan(in, &l, &r)
		if l <= maxL[r] {
			ans[i] = '1'
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { cf1142B(bufio.NewReader(os.Stdin), os.Stdout) }
