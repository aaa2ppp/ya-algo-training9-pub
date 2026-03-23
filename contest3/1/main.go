package main

import (
	"io"
	"os"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]string) string

func solve(words []string) string {
	var ans strings.Builder
	for _, w := range words {
		m := leftN(w)
		n := rightN(w)
		w = w[m : len(w)-n]
		w = w[m : len(w)-n]
		ans.WriteString(w)
	}
	return ans.String()
}

func leftN(s string) int {
	i := 0
	for ; i < len(s) && s[i] == '\''; i++ {
	}
	return i
}

func rightN(s string) int {
	i := len(s) - 1
	for ; i >= 0 && s[i] == '\''; i-- {
	}
	return len(s) - (i + 1)
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var words []string

	words, _ = ScanWordsLn(br, words)

	ans := solve(words)
	PrintWordLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
