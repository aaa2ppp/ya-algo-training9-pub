package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type person struct{ t, k int }

type solveFunc func(int, []person) int

func solve(d int, p []person) int {
	n := len(p)
	if n == 0 {
		return 1
	}

	t := 0
	pos := 0
	for i := range p {
		if p[i].t-t-d < 0 {
			pos = i + 1
		}
		t += p[i].k
	}

	return pos + 1 // to 1-indexing
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var cases int
	ScanIntLn(br, &cases)

	var n, d int
	var p []person
	for range cases {
		ScanIntLn(br, &n, &d)
		p = Resize(p, n)
		for i := range n {
			ScanIntLn(br, &p[i].t, &p[i].k)
		}
		ans := solve(d, p)
		PrintIntLn(bw, ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
