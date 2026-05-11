package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2200E(in io.Reader, out io.Writer) {
	const mx int = 1e6 + 1
	lpf := [mx]int{1: 1}
	for i := 2; i < mx; i++ {
		if lpf[i] == 0 {
			for j := i; j < mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		b := make([]int, n)
		win := false
		for i := range a {
			Fscan(in, &a[i])
			v := a[i]
			p := lpf[v]
			for v /= p; v > 1 && v%p == 0; v /= p {
			}
			if v > 1 {
				win = true
			} else {
				b[i] = p
			}
		}
		if slices.IsSorted(a) || !win && slices.IsSorted(b) {
			Fprintln(out, "Bob")
		} else {
			Fprintln(out, "Alice")
		}
	}
}

//func main() { cf2200E(bufio.NewReader(os.Stdin), os.Stdout) }
