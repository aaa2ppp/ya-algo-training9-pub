package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) bool

func solve(a []int) bool {
	var s Stack[int]
	next := 1
	for _, v := range a {
		s.Push(v)
		for !s.Empty() && s.Top() == next {
			s.Pop()
			next++
		}
	}
	return s.Empty()
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	a := make([]int, n)
	ScanInts(br, a)

	if solve(a) {
		PrintWordLn(bw, "YES")
	} else {
		PrintWordLn(bw, "NO")
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
