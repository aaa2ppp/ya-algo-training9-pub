package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(s, p []int) (int, int)

func solve(s, p []int) (int, int) {
	n, m := len(s), len(p)
	minimum := abs(s[0] - p[0])
	minI, minJ := 0, 0

	i, j := 0, 0
	for i < n && j < m {
		sub := abs(s[i] - p[j])
		if sub < minimum {
			minimum = sub
			minI = i
			minJ = j
		}
		if s[i] < p[j] {
			i++
		} else if s[i] > p[j] {
			j++
		} else {
			minimum = 0
			minI = i
			minJ = j
			break
		}
	}
	return s[minI], p[minJ]
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

	var n, m int
	var shirts, pants []int

	ScanInt(br, &n)

	shirts = Resize(shirts, n)
	ScanInts(br, shirts)

	ScanInt(br, &m)

	pants = Resize(pants, m)
	ScanInts(br, pants)

	a, b := solve(shirts, pants)
	PrintIntLn(bw, a, b)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
