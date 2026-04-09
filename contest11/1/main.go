package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]byte) bool

func solve(a []byte) bool {
	if len(a) == 0 {
		return true
	}
	buf := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		n := copy(buf, a[i:])
		copy(buf[n:], a[:i])
		if check(buf) {
			return true
		}
	}
	return false
}

func check(a []byte) bool {
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

	a, _ := ScanBytes(br, '\n')

	if solve(a) {
		bw.WriteString("YES\n")
	} else {
		bw.WriteString("NO\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
