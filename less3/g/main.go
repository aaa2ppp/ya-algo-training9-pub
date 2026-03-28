package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int32, int32) int64

func solve(d []int32, critR int32) int64 {
	n := len(d)
	var l, r int
	var ans int64 // max=300000^2
	for r < n {
		if d[r]-d[l] <= critR {
			r++
			continue
		}
		ans += int64(n - r)
		l++
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, r int
	var d []int32

	ScanInt(br, &n, &r)

	d = Resize(d, n)
	ScanInts(br, d)

	ans := solve(d, int32(r))
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
