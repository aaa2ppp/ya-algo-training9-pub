package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) byte

func solve(s string) byte {
	n := len(s)
	sum := 0
	minScore := 26
	for _, c := range []byte(s) {
		score := int(25 - (c - 'A'))
		minScore = min(minScore, score)
		sum += score
	}
	mid := ((sum * 2) + n) / (n * 2)
	if mid > minScore+1 {
		mid = minScore + 1
	}
	return byte(25-mid) + 'A'
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')

	ans := solve(s)
	bw.WriteByte(ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
