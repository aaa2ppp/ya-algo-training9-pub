package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(int, []int) []int

func solve(k int, a []int) []int {
	n := len(a)

	var q Deque[int]
	for i := 0; i < k; i++ {
		for q.Len() > 0 && q.Back() > a[i] {
			q.PopBack()
		}
		q.PushBack(a[i])
	}

	ans := make([]int, 0, n-k+1)
	ans = append(ans, q.Front())

	for l, r := 0, k; r < n; l, r = l+1, r+1 {
		if q.Front() == a[l] {
			q.PopFront()
		}
		for q.Len() > 0 && q.Back() > a[r] {
			q.PopBack()
		}
		q.PushBack(a[r])
		ans = append(ans, q.Front())
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, k int
	ScanInt(br, &n, &k)

	a := make([]int, n)
	ScanInts(br, a)

	ans := solve(k, a)
	op := WO{Sep: "\n", End: "\n"}
	PrintInts(bw, op, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
