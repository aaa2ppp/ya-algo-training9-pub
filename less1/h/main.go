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

	var a, b, c int
	ScanInt(br, &a, &b, &c)

	// sqrt(a*x + b) == c

	if c < 0 {
		bw.WriteString("NO SOLUTION\n")
		return
	}

	if a == 0 {
		if b == c*c {
			bw.WriteString("MANY SOLUTIONS\n")
			return
		}
		bw.WriteString("NO SOLUTION\n")
		return
	}

	x := (c*c - b) / a
	if x*a+b != c*c {
		bw.WriteString("NO SOLUTION\n")
		return
	}

	PrintIntLn(bw, x)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
