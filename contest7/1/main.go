package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) int

func solve(s string) int {
	n := len(s)
	m := make(map[int]int)
	b := make([]int, n)

	balance := 0
	for i, c := range []byte(s) {
		switch c {
		case 'a':
			balance++
		case 'b':
			balance--
		}
		m[balance]++
		b[i] = balance
	}

	balance = 0
	count := 0
	for i := range []byte(s) {
		count += m[balance]
		balance = b[i]
		m[balance]--
	}

	return count
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	s, _ := ScanString(br, '\n')
	s = s[:n]

	ans := solve(s)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
