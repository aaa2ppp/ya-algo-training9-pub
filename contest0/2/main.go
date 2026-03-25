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
	if n == 1 {
		return []int{1}
	}

	s := make([]int, n)
	s[0] = a[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + a[i]
	}

	ans := make([]int, n)
	if a[n-1] > a[0] {
		ans[n-1] = 1
		for i := n - 2; i > 0 && a[i] > a[0] && s[i] > a[i+1]; i-- {
			ans[i] = 1
		}
	}

	return ans
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

	op := WO{Sep: "\n", End: "\n"}
	PrintInts(bw, op, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
