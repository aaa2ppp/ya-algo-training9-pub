package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type event struct {
	l, r, x int
}

type solveFunc func([]event, []int) []int

func solve(e []event, q []int) []int {
	ans := make([]int, 0, len(q))
	for _, k := range q {
		s := 0
		for i := 1; i < len(e); i++ {
			if !(e[i].l <= k && k <= e[i].r) {
				continue
			}
			if (k-e[i].l)%2 == 0 {
				s += e[i].x
			} else {
				s -= e[i].x
			}
		}
		ans = append(ans, s)
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, m int
	ScanIntLn(br, &n, &m)

	e := make([]event, n+1)
	for i := 1; i <= n; i++ {
		var l, r, x int
		ScanIntLn(br, &l, &r, &x)
		e[i] = event{l, r, x}
	}

	q := make([]int, m)
	ScanInts(br, q)

	ans := solve(e, q)

	op := WO{Sep: "\n", End: "\n"}
	PrintInts(bw, op, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
