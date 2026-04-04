package main

import (
	"io"
	"os"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(int, string, string) string

func solve(n int, order, prefix string) string {
	var ans strings.Builder
	var st Stack[byte]

	const push = 1
	m1 := []byte{
		'[': push,
		'(': push,
		']': '[',
		')': '(',
	}
	m2 := []byte{
		'[': ']',
		'(': ')',
	}

	for _, c := range []byte(prefix) {
		ans.WriteByte(c)
		if m1[c] == push {
			st.Push(c)
		} else {
			// нам гарантируют корректный ввод
			st.Pop()
		}
	}

	n -= ans.Len()
	for ; st.Len() < n; n-- {
		for _, c := range []byte(order) {
			if m1[c] == push {
				st.Push(c)
				ans.WriteByte(c)
				break
			}
			if !st.Empty() && m1[c] == st.Top() {
				st.Pop()
				ans.WriteByte(c)
				break
			}
		}
	}

	for ; n > 0; n-- {
		c := st.Pop()
		c = m2[c]
		ans.WriteByte(c)
	}

	return ans.String()
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	order, _ := ScanString(br, '\n')
	prefix, _ := ScanString(br, '\n')

	ans := solve(n, order, prefix)
	bw.WriteString(ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
