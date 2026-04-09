package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string, []string) [][2]int

func solve(a string, b []string) [][2]int {
	var ans [][2]int
	for i := range b {
		for j := i + 1; j < len(b); j++ {
			if similar(a, b[i], b[j]) {
				ans = append(ans, [2]int{i + 1, j + 1}) // to 1-indexing
			}
		}
	}
	return ans
}

func similar(a, b1, b2 string) bool {
	n := len(a)
	var s1, s2 int
	var v1, v2 int
	for i := 0; i < n; i++ {
		if b1[i] == a[i] {
			s1++
		}
		if b2[i] == a[i] {
			s2++
		}
		if b1[i] == b2[i] {
			if b1[i] == a[i] {
				v1++
			} else {
				v2++
			}
		}
	}
	return v1 > s1/2 && v1 > s2/2 && v2 > (n-s1)/2 && v2 > (n-s2)/2
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	a, _ := ScanString(br, '\n')

	var m int
	ScanIntLn(br, &m)

	b := make([]string, m)
	ScanWords(br, b)

	ans := solve(a, b)
	PrintIntLn(bw, len(ans))
	for _, v := range ans {
		PrintIntLn(bw, v[0], v[1])
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
