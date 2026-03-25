package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(p, v, m, q int) int

func solve(p, v, m, q int) int {
	a, b := p-v, p+v
	c, d := m-q, m+q

	maxBeg := max(a, c)
	minEnd := min(b, d)

	dots := (b - a + 1) + (d - c + 1)

	if maxBeg <= minEnd { // пересекаются
		dots -= minEnd - maxBeg + 1
	}

	return dots
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var p, v, m, q int
	ScanInt(br, &p, &v, &m, &q)

	ans := solve(p, v, m, q)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
