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
	ans := make([]int, n)

	type item struct{ idx, val int }
	var st Stack[item]

	for i, val := range a {
		for st.Len() > 0 && val < st.Top().val {
			it := st.Pop()
			ans[it.idx] = i
		}
		ans[i] = -1
		st.Push(item{i, val})
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	a := make([]int, n)
	ScanInts(br, a)

	a = solve(a)
	PrintIntsLn(bw, a)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
