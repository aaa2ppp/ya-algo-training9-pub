package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(a []int) int {
	n := len(a)

	forbidden := make([]bool, n)
	opposite := make([]bool, n)

	for i, v := range a {
		v-- // to 0-indexing
		forbidden[(v+n-i)%n] = true
		opposite[(i+n-v)%n] = true
	}

	for k := 0; k < n; k++ {
		// если лень думать, в каком направлении нужно вращать,
		// можно просто проверить на примере.
		if !forbidden[k] {
			return k
		}
		// if !opposite[k] {
		// 	return k
		// }
	}

	return -1
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	a := make([]int, n)
	ScanInts(br, a)

	k := solve(a)
	PrintIntLn(bw, k)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
