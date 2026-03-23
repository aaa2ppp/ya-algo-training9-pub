package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(tRoom, tCond int, mode string) int

func solve(tRoom, tCond int, mode string) int {
	var t int
	switch mode {
	case "freeze":
		t = min(tRoom, tCond)
	case "heat":
		t = max(tRoom, tCond)
	case "auto":
		t = tCond
	default: // "fan"
		t = tRoom
	}
	return t
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var tRoom, tCond int
	var mode string

	ScanIntLn(br, &tRoom, &tCond)
	ScanWordLn(br, &mode)

	t := solve(tRoom, tCond, mode)
	PrintIntLn(bw, t)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
