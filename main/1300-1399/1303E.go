package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1303E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s, t []byte
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		for i := range t {
			t[i] -= 'a'
		}
		n := len(s)
		r := [26]int{}
		for i := range r {
			r[i] = n
		}
		nxt := make([][26]int, n+2)
		for i := range nxt[n+1] {
			nxt[n+1][i] = n
		}
		nxt[n] = r
		for i := n - 1; i >= 0; i-- {
			r[s[i]-'a'] = i
			nxt[i] = r
		}

		m := len(t)
		f := make([][]int, m)
		for i := range f {
			f[i] = make([]int, m)
		}
		for k := 1; k < m; k++ {
			for i, v := range t[:k] {
				for j, w := range t[k:] {
					if i == 0 && j == 0 {
						f[i][j] = -1
					} else if i == 0 {
						f[i][j] = nxt[f[i][j-1]+1][w]
					} else if j == 0 {
						f[i][j] = nxt[f[i-1][j]+1][v]
					} else {
						f[i][j] = min(nxt[f[i-1][j]+1][v], nxt[f[i][j-1]+1][w])
					}
				}
			}
			if f[k-1][n-k-1] < n {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { CF1303E(os.Stdin, os.Stdout) }
