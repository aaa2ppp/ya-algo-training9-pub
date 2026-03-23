package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]string) int

func solve(mx []string) int {
	n, m := len(mx), len(mx[0])
	var ans int

	for i := 0; i < n; i++ {
		prev := mx[i][0]
		for j := 1; j < m; j++ {
			cur := mx[i][j]
			if cur == '.' && prev == '.' {
				ans++
			}
			prev = cur
		}
	}

	for j := 0; j < m; j++ {
		prev := mx[0][j]
		for i := 1; i < n; i++ {
			cur := mx[i][j]
			if cur == '.' && prev == '.' {
				ans++
			}
			prev = cur
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

	mx := make([]string, n)
	for i := 0; i < n; i++ {
		s, _ := ScanString(br, '\n')
		mx[i] = s
	}

	ans := solve(mx)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
