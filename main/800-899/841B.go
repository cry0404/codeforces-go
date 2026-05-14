package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf841B(in io.Reader, out io.Writer) {
	var n, v, s, cnt1 int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		s += v
		cnt1 += v & 1
	}
	if s&1 > 0 || cnt1 > 0 {
		Fprint(out, "First")
	} else {
		Fprint(out, "Second")
	}
}

//func main() { cf841B(bufio.NewReader(os.Stdin), os.Stdout) }
