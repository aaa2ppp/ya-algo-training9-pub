package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(h []int) int {
	n := len(h)
	type item struct{ idx, h int }
	var st Stack[item]
	ans := 0
	for i := range h {
		last := item{i, h[i]}
		for !st.Empty() && h[i] <= st.Top().h {
			last = st.Pop()
			ans = max(ans, (i-last.idx)*last.h)
			last.h = h[i]
		}
		st.Push(last)
	}
	for !st.Empty() {
		last := st.Pop()
		ans = max(ans, (n-last.idx)*last.h)
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	h := make([]int, n)
	ScanInts(br, h)

	ans := solve(h)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
