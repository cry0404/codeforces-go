package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1743C(in io.Reader, out io.Writer) {
	var T, n, v int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans, mn := 0, 0
		for i, b := range s {
			Fscan(in, &v)
			if b == '1' {
				ans += v
				mn = min(mn, v)
				if i == n-1 || s[i+1] == '0' {
					ans -= mn
				}
			} else if i < n-1 && s[i+1] == '1' {
				ans += v
				mn = v
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1743C(bufio.NewReader(os.Stdin), os.Stdout) }
