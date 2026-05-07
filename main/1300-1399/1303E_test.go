package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1303/problem/E
// https://codeforces.com/problemset/status/1303/problem/E
func TestCF1303E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
ababcd
abcba
a
b
defi
fed
xyz
x
outputCopy
YES
NO
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1303E)
}
