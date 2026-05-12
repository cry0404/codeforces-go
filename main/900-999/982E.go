package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf982E(in io.Reader, out io.Writer) {
	var n, m, x, y, vx, vy int
	Fscan(in, &n, &m, &x, &y, &vx, &vy)

	if vx == 0 {
		if 0 < x && x < n {
			Fprint(out, -1)
		} else if vy < 0 {
			Fprint(out, x, 0)
		} else {
			Fprint(out, x, m)
		}
		return
	}

	if vy == 0 {
		if 0 < y && y < m {
			Fprint(out, -1)
		} else if vx < 0 {
			Fprint(out, 0, y)
		} else {
			Fprint(out, n, y)
		}
		return
	}

	k1 := x % n
	if vx > 0 {
		k1 = (n - x) % n
	}

	k2 := y % m
	if vy > 0 {
		k2 = (m - y) % m
	}

	g := gcd82(n, m)
	if (k1-k2)%g != 0 {
		Fprint(out, -1)
		return
	}

	lcm := n / g * m
	_, xx, _ := exgcd82(n, m)
	xx *= (k2 - k1) / g
	xx %= m / g
	s := ((xx*n+k1)%lcm + lcm) % lcm
	ansX := x + s*vx
	ansY := y + s*vy
	n *= 2
	m *= 2
	Fprint(out, (ansX%n+n)%n, (ansY%m+m)%m)
}

//func main() { cf982E(os.Stdin, os.Stdout) }

func exgcd82(a, b int) (g, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, y, x = exgcd82(b, a%b)
	y -= a / b * x
	return
}

func gcd82(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
