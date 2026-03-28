package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string, int) (int, int)

func solve(s string, k int) (int, int) {
	n := len(s)
	freq := make([]int, 26)
	var max_d, max_i int

	for l, r := 0, -1; r < n && l < n; {
		r++
		if r >= n {
			break
		}
		i := s[r] - 'a'
		freq[i]++
		if freq[i] > k {
			for ; l < r && s[l] != s[r]; l++ {
				freq[s[l]-'a']--
			}
			freq[s[l]-'a']--
			l++
		}
		if d := r - l + 1; d > max_d {
			max_d = d
			max_i = l
		}
	}

	return max_d, max_i
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, k int
	ScanIntLn(br, &n, &k)

	s, _ := ScanString(br, '\n')
	s = s[:n]

	l, i := solve(s, k)
	PrintIntLn(bw, l, i+1) // to 1-indexing
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
