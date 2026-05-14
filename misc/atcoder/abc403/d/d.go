package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, d, v int
	Fscan(in, &n, &d)
	if d == 0 {
		d = 1e9
	}
	g := map[int][]int{}
	for range n {
		Fscan(in, &v)
		g[v%d] = append(g[v%d], v)
	}
	if d == 1e9 {
		Fprint(out, n-len(g))
		return
	}

	ans := n
	for _, a := range g {
		slices.Sort(a)
		f0, f1 := 0, 0
		cnt := 0
		for i, v := range a {
			cnt++
			if i == len(a)-1 || a[i+1]-v == d {
				f0, f1 = f1, max(f1, f0+cnt)
				cnt = 0
			} else if a[i+1] > v {
				ans -= max(f1, f0+cnt)
				f0, f1 = 0, 0
				cnt = 0
			}
		}
		ans -= f1
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
