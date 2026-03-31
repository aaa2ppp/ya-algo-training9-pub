package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(p []int) int {
	n := len(p)
	m := make(map[int]int, n)
	l, r := 0, 0
	t := p[0]
	m[t]++
	count := 1
	ans := 0
	for {
		if count == 2 {
			ans = max(ans, r-l+1)
		}
		if count <= 2 {
			r++
			if r >= len(p) {
				break
			}
			t := p[r]
			if m[t] == 0 {
				count++
			}
			m[t]++
		} else {
			t := p[l]
			m[t]--
			if m[t] == 0 {
				count--
			}
			l++
			if l >= len(p) {
				break
			}
		}
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	p := make([]int, n)
	ScanInts(br, p)

	ans := solve(p)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
