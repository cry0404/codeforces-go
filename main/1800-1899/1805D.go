package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1805D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	tmp := make([]int, n)
	maxD, u := -1, 0
	var dfs func(int, int, int)
	dfs = func(v, fa, d int) {
		tmp[v] = d
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
			}
		}
	}

	dfs(0, -1, 0)

	maxD = -1
	dfs(u, -1, 0)
	d1 := slices.Clone(tmp)

	maxD = -1
	dfs(u, -1, 0)
	d2 := tmp

	s := make([]int, n+1)
	for i, d := range d1 {
		if i != u { // u 所在的连通块算在下面的 sum := 1 中
			// 当 k > max(d1[i], d2[i]) 时，i 是孤立点
			s[max(d, d2[i])+1]++ 
		}
	}

	sum := 1
	for i := 1; i <= n; i++ {
		sum += s[i]
		Fprint(out, sum, " ")
	}
}

//func main() { cf1805D(bufio.NewReader(os.Stdin), os.Stdout) }
