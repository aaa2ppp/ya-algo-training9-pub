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
	curLen := 0
	maxLen := 0
	var prev byte
	for i := 0; i < n; i++ {
		cur := s[i]
		if prev == 'a' && cur == 'h' || prev == 'h' && cur == 'a' {
			curLen++
			maxLen = max(maxLen, curLen)
		} else {
			curLen = 0
		}
		if curLen == 0 && (cur == 'a' || cur == 'h') {
			curLen = 1
			maxLen = max(maxLen, curLen)
		}
		prev = cur
	}
	return maxLen
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
