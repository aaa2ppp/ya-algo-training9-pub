package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int64) int64

func solve(a []int64) int64 {
	n := len(a)
	b := make([]int64, n+1)
	for i, v := range a {
		b[i+1] = b[i] + v
	}

	totalMax := b[1]

	l, r := 0, 1
	for r <= n {
		curMax := b[r] - b[l]
		totalMax = max(totalMax, curMax)
		if curMax < 0 {
			l = r
		}
		r++
	}

	return totalMax
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var a []int64

	ScanIntLn(br, &n)

	a = Resize(a, n)
	ScanInts(br, a)

	ans := solve(a)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
