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
		var ans, sum, mn int
		for i, b := range s {
			Fscan(in, &v)
			if b == '1' {
				sum += v
				mn = min(mn, v)
				if i == n-1 || s[i+1] == '0' {
					ans += sum - mn
				}
			} else if i < n-1 && s[i+1] == '1' {
				sum, mn = v, v
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1743C(bufio.NewReader(os.Stdin), os.Stdout) }
