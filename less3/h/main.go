package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) (int, int, int)

func solve(a []int) (int, int, int) {
	n := len(a)

	min_sub := int(1e15) + 1
	min_l, min_r := -1, -1

	sv, sm := a[0], a[n-1]
	for l, r := 0, n-1; l < r; {
		sub := abs(sv - sm)
		if sub < min_sub {
			min_sub = sub
			min_l = l
			min_r = r
		}
		if sv < sm {
			l++
			if l >= r {
				break
			}
			sv += a[l]
		} else if sv > sm {
			r--
			if l >= r {
				break
			}
			sm += a[r]
		} else {
			return 0, l, r
		}

	}

	return min_sub, min_l, min_r
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	a := make([]int, n)
	ScanInts(br, a)

	s, l, r := solve(a)
	PrintIntLn(bw, s, l+1, r+1) // to 1-indexing
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
