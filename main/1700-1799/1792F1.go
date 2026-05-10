package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1792F1(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n int
	Fscan(in, &n)
	f := make([]int, n+1)
	z := make([]int, n+1)
	c := make([]int, n+1)
	f[1] = 1
	z[0] = 1
	z[1] = 1
	c[0] = 1
	c[1] = 1
	for i := 2; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			c[j] = (c[j] + c[j-1]) % mod
			z[i] = (z[i] + f[j]*z[i-j]%mod*c[j-1]) % mod
		}
		c[i] = 1
		f[i] = z[i]
		z[i] = z[i] * 2 % mod
	}
	Fprint(out, (f[n]*2-2+mod)%mod)
}

//func main() { cf1792F1(os.Stdin, os.Stdout) }
