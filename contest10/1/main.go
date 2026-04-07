package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([][]byte) int

func solve(mx [][]byte) int {
	n, m := len(mx), len(mx[0])

	fill := func(i0, j0, d int) bool {
		i1 := i0 + d
		if i1 > n {
			return false
		}
		j1 := j0 + d
		if j1 > m {
			return false
		}
		for i := i0; i < i1; i++ {
			if mx[i][j1-1] != '#' {
				return false
			}
		}
		for j := j0; j < j1; j++ {
			if mx[i1-1][j] != '#' {
				return false
			}
		}
		for i := i0; i < i1; i++ {
			mx[i][j1-1] = '.'
		}
		for j := j0; j < j1; j++ {
			mx[i1-1][j] = '.'
		}
		return true
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mx[i][j] == '#' {
				ans++
				for d := 1; fill(i, j, d); d++ {
				}
			}
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, m int
	ScanIntLn(br, &n, &m)

	mx := make([][]byte, n)
	for i := range mx {
		mx[i], _ = ScanBytes(br, '\n')
		mx[i] = mx[i][:m]
	}

	ans := solve(mx)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
