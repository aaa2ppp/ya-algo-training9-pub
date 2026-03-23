package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func()

func solve() {}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int

	ScanIntLn(br, &n)

	syns := make(map[string]string, n*2)
	for i := 0; i < n; i++ {
		var w1, w2 string
		ScanWordLn(br, &w1, &w2)
		syns[w1] = w2
		syns[w2] = w1
	}

	var target string
	ScanWordLn(br, &target)

	PrintWordLn(bw, syns[target])
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
