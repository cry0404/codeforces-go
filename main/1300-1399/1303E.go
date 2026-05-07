package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1303E(in io.Reader, out io.Writer) {
	var T int
	var s, t []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		f := make([]int, len(t))
		for mid := range t {
			a, b := t[:mid], t[mid:]
			f[0] = 0
			for j := 1; j <= mid; j++ {
				f[j] = -1e9
			}
			for _, x := range s {
				for j := mid; j >= 0; j-- {
					k := f[j]
					res := k
					if 0 <= k && k < len(b) && x == b[k] {
						res = k + 1
					}
					if j > 0 && x == a[j-1] {
						res = max(res, f[j-1])
					}
					f[j] = res
				}
			}
			if f[mid] == len(b) {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1303E(os.Stdin, os.Stdout) }
