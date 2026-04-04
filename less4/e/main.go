package main

import (
	"io"
	"os"
	"strconv"
	"unicode"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]string) int

func solve(expr []string) int {
	var st Stack[int]
	for _, token := range expr {
		c := token[0]
		switch {
		case unicode.IsDigit(rune(c)):
			v, _ := strconv.Atoi(token)
			st.Push(v)
		case c == '+':
			b, a := st.Pop(), st.Pop()
			st.Push(a + b)
		case c == '-':
			b, a := st.Pop(), st.Pop()
			st.Push(a - b)
		case c == '*':
			b, a := st.Pop(), st.Pop()
			st.Push(a * b)
		}
	}
	return st.Pop()
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var expr []string
	expr, _ = ScanWordsLn(br, expr)

	ans := solve(expr)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
