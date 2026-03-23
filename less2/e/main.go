package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type statement struct {
	a, b int
}

type solveFunc func([]statement) int

func solve(st []statement) int {
	n := len(st)
	set := make(map[statement]bool, n)

	for _, k := range st {
		set[k] = true
	}
	var ans int
	for i := 0; i < n; i++ {
		if set[statement{i, n - i - 1}] {
			ans++
		}
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var st []statement

	ScanIntLn(br, &n)
	st = Grow(st, n)
	for i := 0; i < n; i++ {
		var a, b int
		ScanIntLn(br, &a, &b)
		st = append(st, statement{a, b})
	}

	ans := solve(st)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
