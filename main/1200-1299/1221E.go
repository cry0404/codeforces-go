package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

func cf1221E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	in1 := func(x, a, b int) bool { return a <= x && x < 2*b }
	in2 := func(x, a, b int) bool { return x >= 2*b || b <= x && x < a }
	f := func(a, b int, s string) bool {
		n := len(s)
		q := []int{}
		for i := 0; i < n; i++ {
			if s[i] == '.' {
				j := i
				for j+1 < n && s[j+1] == '.' {
					j++
				}
				q = append(q, j-i+1)
				i = j
			}
		}
		if len(q) == 0 {
			return false
		}

		slices.Sort(q)
		x, ok := 0, false
		for i := range len(q) - 1 {
			if in1(q[i], a, b) {
				x ^= 1
			}
			if in2(q[i], a, b) {
				ok = true
			}
		}
		for i, j := 0, q[len(q)-1]-a; j >= 0; i, j = i+1, j-1 {
			if !ok && !in2(i, a, b) && !in2(j, a, b) && x^b2i21(in1(i, a, b))^b2i21(in1(j, a, b)) == 0 {
				return true
			}
		}
		return false
	}

	var T, a, b int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &s)
		if f(a, b, s) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1221E(bufio.NewReader(os.Stdin), os.Stdout) }

func b2i21(b bool) int {
	if b {
		return 1
	}
	return 0
}
