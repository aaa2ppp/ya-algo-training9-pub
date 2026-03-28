package main

import (
	"io"
	"iter"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

const modulo = 1000000007

func NewRandSeq(x int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		for {
			if !yield(x) {
				return
			}
			x = (11173*x + 1) % modulo
		}
	}
}

type solveFunc func(a []int, q int, seq iter.Seq[int]) int

func solve(a []int, q int, seq iter.Seq[int]) int {
	n := len(a)
	b := make([]int, n+1)
	for i, v := range a {
		b[i+1] = (b[i] + v) % modulo
	}

	nextX, stop := iter.Pull(seq)
	defer stop()

	next := func() (int, int) {
		x0, _ := nextX()
		x1, _ := nextX()
		x0 %= n
		x1 %= n
		if x0 > x1 {
			return x1, x0
		}
		return x0, x1
	}

	var ans int
	for i := 0; i < q; i++ {
		l, r := next()
		ans = (ans + (modulo - b[l])) % modulo
		ans = (ans + b[r+1]) % modulo
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, q, x0 int
	var a []int

	ScanIntLn(br, &n)

	a = Resize(a, n)
	ScanInts(br, a)

	ScanInt(br, &q, &x0)
	rand := NewRandSeq(x0)

	ans := solve(a, q, rand)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
