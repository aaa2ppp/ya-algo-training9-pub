package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(a, b []int) (int, int)

func solve(a, b []int) (int, int) {
	const maxN = 1e6

	aq := NewDequeFrom(a)
	bq := NewDequeFrom(b)

	n := 0
	for ; n < maxN && !aq.Empty() && !bq.Empty(); n++ {
		ac := aq.PopFront()
		bc := bq.PopFront()
		win := aq
		switch {
		case ac == 0 && bc == 9:
		case ac == 9 && bc == 0:
			win = bq
		case ac < bc:
			win = bq
		}
		win.PushBack(ac)
		win.PushBack(bc)
	}

	switch {
	case aq.Empty():
		return 2, n
	case bq.Empty():
		return 1, n
	}
	return 0, -1
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	a := make([]int, 5)
	b := make([]int, 5)
	ScanInts(br, a)
	ScanInts(br, b)

	winner, n := solve(a, b)
	names := []string{"botva", "first", "second"}
	PrintAnyLn(bw, names[winner], n)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
