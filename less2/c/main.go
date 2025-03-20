package main

import (
	"io"
	"os"
	"slices"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(a, b []int) []int

func solve(a, b []int) []int {
	if len(a) > len(b) {
		a, b = b, a
	}
	set := make(map[int]bool, len(a))
	ans := make([]int, 0, len(a))

	for _, v := range a {
		set[v] = true
	}
	for _, v := range b {
		if set[v] {
			ans = append(ans, v)
		}
	}
	slices.Sort(ans)

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var a, b []int

	a, _ = ScanIntsLn(br, a)
	b, _ = ScanIntsLn(br, b)

	ans := solve(a, b)
	PrintIntsLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
