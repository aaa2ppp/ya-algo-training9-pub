package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int, int) int

func solve(a []int, k int) int {
	n := len(a)
	b := make([]int, n+1)
	for i, v := range a {
		b[i+1] = b[i] + v
	}
	var ans int
	l, r := 0, 1
	for r <= n {
		s := b[r] - b[l]
		if s < k {
			r++
		} else if s > k {
			l++
		} else {
			o, p := 1, 1
			for r < n && a[r] == 0 {
				r++
				p++
			}
			l++
			for l < r && a[l] == 0 {
				l++
				o++
			}
			ans += o * p
		}
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, k int
	var a []int

	ScanIntLn(br, &n, &k)

	a = Resize(a, n)
	ScanInts(br, a)

	ans := solve(a, k)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
