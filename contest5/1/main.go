package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type signal struct {
	x, d int
}

type solveFunc func([]signal) int

func solve(s []signal) int {
	min_x, max_x := s[0].x-s[0].d, s[0].x+s[0].d

	for i := range s {
		min_x = max(min_x, s[i].x-s[i].d)
		max_x = min(max_x, s[i].x+s[i].d)
	}

	if min_x > max_x {
		return -1
	}
	return max_x
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	s := make([]signal, n)
	for i := 0; i < n; i++ {
		ScanInt(br, &s[i].x, &s[i].d)
	}

	ans := solve(s)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
