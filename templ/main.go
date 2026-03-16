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
	_ = n
	return nil
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var a []int

	if _, err := ScanIntLn(br, &n); err != nil {
		panic(err)
	}

	a = Resize(a, n)
	if _, err := ScanInts(br, a); err != nil {
		panic(err)
	}

	ans := solve(a)
	PrintIntsLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
