package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int, int) (int, int)

func solve(a []int, k int) (int, int) {
	n := len(a)
	freq := make([]int, k+1)
	freq[a[0]] = 1
	min_d, min_l, min_r := n+1, -1, -1
	count := 1

	for l, r := 0, 0; l+k <= n && r < n; {
		if count < k {
			r++
			if r >= n {
				break
			}
			i := a[r]
			if freq[i] == 0 {
				count++
			}
			freq[i]++
			continue
		}
		if d := r - l + 1; d < min_d {
			min_d = d
			min_l = l
			min_r = r
		}
		i := a[l]
		freq[i]--
		if freq[i] == 0 {
			count--
		}
		l++
		if l+k > n {
			break
		}
	}

	return min_l, min_r
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, k int
	ScanInt(br, &n, &k)

	a := make([]int, n)
	ScanInts(br, a)

	l, r := solve(a, k)
	PrintIntLn(bw, l+1, r+1) // to 1-indexing
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
