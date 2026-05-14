package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1980F2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		type tuple struct{ c, r, i int }
		a := make([]tuple, k+1)
		for i := range k {
			var r, c int
			Fscan(in, &r, &c)
			a = append(a, tuple{c, r, i})
		}
		a[k] = tuple{m + 1, n + 1, k}
		slices.SortFunc(a, func(a, b tuple) int { return cmp.Or(a.c-b.c, a.r-b.r) })

		ans := make([]int, k+1)
		sum := 0
		preC := 0
		s := tuple{0, 0, k}
		st := tuple{}
		for _, t := range a {
			c, r := t.c, t.r
			if r < st.c {
				continue
			}
			if r < s.c {
				st.i += (c - st.r) * (st.c - preC)
				st.c = r
				st.r = c
			} else {
				ans[s.i] -= st.i + (c-st.r)*(st.c-preC)
				ans[s.i] += (c - s.r) * (s.c - preC)
				sum += (c - s.r) * s.c
				preC = s.c
				st = tuple{s.c, c, 0}
				s = tuple{r, c, t.i}
			}
		}

		Fprintln(out, m*n-sum)
		for _, v := range ans[:k] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1980F2(bufio.NewReader(os.Stdin), os.Stdout) }
