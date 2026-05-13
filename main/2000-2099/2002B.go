package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2002B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		rev := slices.Clone(b)
		slices.Reverse(rev)
		if slices.Equal(a, b) || slices.Equal(a, rev) {
			Fprintln(out, "Bob")
		} else {
			Fprintln(out, "Alice")
		}
	}
}

//func main() { cf2002B(bufio.NewReader(os.Stdin), os.Stdout) }
