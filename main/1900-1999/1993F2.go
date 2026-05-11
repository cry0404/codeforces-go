package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1993F2(in io.Reader, out io.Writer) {
	var T, N, k, m, n int
	var str string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &N, &k, &m, &n, &str)
		n *= 2
		m *= 2
		px := make([]int, N+1)
		py := make([]int, N+1)
		for i, b := range str {
			addX := 0
			if b == 'D' {
				addX = n - 1
			} else if b == 'U' {
				addX = 1
			}
			px[i+1] = (px[i] + addX) % n

			addY := 0
			if b == 'L' {
				addY = m - 1
			} else if b == 'R' {
				addY = 1
			}
			py[i+1] = (py[i] + addY) % m
		}

		dx := px[N]
		dy := py[N]
		gx := gcd93(dx, n)
		gy := gcd93(dy, m)
		p := n / gx
		q := m / gy
		dx /= gx
		dy /= gy
		ix := inv93(dx, p)
		iy := inv93(dy, q)
		d, u, _ := exgcd93(p, q)
		e := p / d * q

		ans := 0
		for i := 1; i <= N; i++ {
			rx := (n - px[i]) % n
			ry := (m - py[i]) % m
			if rx%gx != 0 || ry%gy != 0 {
				continue
			}

			rx = rx / gx * ix % p
			ry = ry / gy * iy % q
			if (ry-rx)%d != 0 {
				continue
			}

			z := q / d
			s := (((ry-rx)/d*u)%z + z) % z
			k0 := s*p + rx
			if k0 < k {
				ans += (k-k0-1)/e + 1
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1993F2(bufio.NewReader(os.Stdin), os.Stdout) }

func exgcd93(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, y, x = exgcd93(b, a%b)
	y -= a / b * x
	return
}

func inv93(a, m int) int {
	_, x, _ := exgcd93(a, m)
	return (x%m + m) % m
}

func gcd93(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
