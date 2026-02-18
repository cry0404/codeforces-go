package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2040D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 4e5 + 1
	np := [mx]bool{1: true}
	for i := 2; i*i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
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

		ans := make([]int, n)
		ans[0] = 1
		cur := 2
		var dfs func(int, int)
		dfs = func(v, fa int) {
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				for !np[cur-ans[v]] {
					cur++
				}
				ans[w] = cur
				cur++
				dfs(w, v)
			}
		}
		dfs(0, -1)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2040D(bufio.NewReader(os.Stdin), os.Stdout) }
