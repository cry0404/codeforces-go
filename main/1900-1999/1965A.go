package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1965A(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		a = slices.Compact(a)
		pre := 0
		for i, v := range a {
			if i == len(a)-1 || v-pre > 1 {
				if i%2 == 0 {
					Fprintln(out, "Alice")
				} else {
					Fprintln(out, "Bob")
				}
				break
			}
			pre = v
		}
	}
}

//func main() { cf1965A(bufio.NewReader(os.Stdin), os.Stdout) }
