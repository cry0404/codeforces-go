package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2004C(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)

		s := 0
		for i := n - 1; i > 0; i -= 2 {
			s += a[i] - a[i-1]
		}
		Fprintln(out, max(s-k, 0)+n%2*a[0])
	}
}

//func main() { cf2004C(bufio.NewReader(os.Stdin), os.Stdout) }
