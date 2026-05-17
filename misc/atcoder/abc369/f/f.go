package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// https://github.com/EndlessCheng
type pair struct{ v, i int }

func max(a, b pair) pair {
	if a.v < b.v {
		return b
	}
	return a
}

type fenwick []pair

func (f fenwick) update(i int, p pair) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], p)
	}
}

func (f fenwick) pre(i int) (res pair) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return
}

func run(in io.Reader, out io.Writer) {
	var h, w, n int
	Fscan(in, &h, &w, &n)
	type point struct{ r, c int }
	a := make([]point, n+2)
	for i := range n {
		Fscan(in, &a[i].r, &a[i].c)
	}
	a[n] = point{1, 1}
	a[n+1] = point{h, w}
	slices.SortFunc(a, func(a, b point) int { return cmp.Or(a.r-b.r, a.c-b.c) })

	t := make(fenwick, w+1)
	f := make([]int, n+2)
	from := make([]int, n+2)
	for i, p := range a {
		mx := t.pre(p.c)
		f[i] = mx.v + 1
		from[i] = mx.i
		t.update(p.c, pair{f[i], i})
	}

	Fprintln(out, f[n+1]-2)
	ans := make([]byte, 0, h+w-2)
	for i := n + 1; i > 0; {
		j := from[i]
		ans = append(ans, strings.Repeat("R", a[i].c-a[j].c)...)
		ans = append(ans, strings.Repeat("D", a[i].r-a[j].r)...)
		i = j
	}
	slices.Reverse(ans)
	Fprintf(out, "%s", ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
