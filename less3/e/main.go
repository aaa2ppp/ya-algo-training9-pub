package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

const modulo = 1000000007

type solveFunc func([]uint64) uint64

func solve(a []uint64) uint64 {
	n := len(a)
	b := make([]uint64, n+1)
	for i, v := range a {
		b[i+1] = (b[i] + v) % modulo
	}

	var ans uint64
	for i := 1; i < n-1; i++ {
		v := a[i]
		v = (v * (b[i] + (modulo - b[0]))) % modulo
		v = (v * (b[n] + (modulo - b[i+1]))) % modulo
		ans = (ans + v) % modulo
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var a []uint64

	ScanIntLn(br, &n)

	a = Resize(a, n)
	ScanInts(br, a)

	ans := solve(a)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
