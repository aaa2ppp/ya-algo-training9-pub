package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(m, n, h, w int) int

func solve(m, n, h, w int) int {
	var p1, p2, q1, q2 int

	for h<<p1 < m {
		p1++
	}
	for w<<q1 < n {
		q1++
	}

	for w<<p2 < m {
		p2++
	}
	for h<<q2 < n {
		q2++
	}

	return min(p1+q1, p2+q2)
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var m, n, h, w int
	ScanIntLn(br, &m, &n, &h, &w)

	ans := solve(m, n, h, w)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
