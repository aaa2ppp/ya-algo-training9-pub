package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) bool

func solve(a string) bool {
	const push = 1
	m := []byte{
		'(': push,
		'[': push,
		'{': push,
		')': '(',
		']': '[',
		'}': '{',
	}

	var st Stack[byte]

	for _, c := range []byte(a) {
		if m[c] == push {
			st.Push(c)
			continue
		}
		if st.Empty() || st.Top() != m[c] {
			return false
		}
		st.Pop()
	}

	return st.Empty()
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	a, _ := ScanString(br, '\n')

	if solve(a) {
		bw.WriteString("yes\n")
	} else {
		bw.WriteString("no\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
