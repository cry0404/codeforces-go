package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1658E(in io.Reader, out io.Writer) {
	var n, k, v int
	Fscan(in, &n, &k)
	type pair struct{ x, y int }
	p := make([]pair, n*n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			Fscan(in, &v)
			p[v] = pair{i, j}
		}
	}

	ans := make([][]byte, n)
	for i := range ans {
		ans[i] = bytes.Repeat([]byte{'G'}, n)
	}
	var a, b, c, d int = -1e9, -1e9, -1e9, -1e9
	for i := n * n; i > 0; i-- {
		x, y := p[i].x, p[i].y
		if a <= x+y+k && b <= x-y+k && c <= -x+y+k && d <= -x-y+k {
			a = max(a, x+y)
			b = max(b, x-y)
			c = max(c, -x+y)
			d = max(d, -x-y)
			ans[x-1][y-1] = 'M'
		}
	}
	for _, r := range ans {
		Fprintf(out, "%s\n", r)
	}
}

//func main() { cf1658E(bufio.NewReader(os.Stdin), os.Stdout) }
