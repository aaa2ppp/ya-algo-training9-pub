package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) []int

func solve(a []int) []int {
	n := len(a)
	b := make([]int, n+1)
	for i, v := range a {
		b[i+1] = b[i] + v
	}
	return b[1:]
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var a []int

	ScanIntLn(br, &n)

	a = Resize(a, n)
	ScanInts(br, a)

	ans := solve(a)
	PrintIntsLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
