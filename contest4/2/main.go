package main

import (
	"io"
	"os"
	"sort"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]string) int

func solve(words []string) int {
	m := len(words[0]) // все слова одинаковые
	i := sort.Search(m, func(k int) bool {
		return !check(k+1, words)
	})
	return i
}

func check(k int, words []string) bool {
	if k == 0 {
		return true
	}
	m := make(map[string]int)
	for _, w := range words {
		m[w[:k]]++
	}
	var ans int
	for _, v := range m {
		ans += v / 2
	}
	return ans*2 == len(words)
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var words []string

	ScanIntLn(br, &n)

	words = Resize(words, n)
	for i := 0; i < n; i++ {
		w, _ := ScanString(br, '\n')
		words[i] = w
	}

	ans := solve(words)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
