package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) []byte

func solve(freq []int) []byte {
	var ans []byte

	mid := -1
	for i, v := range freq {
		if mid == -1 && v%2 == 1 { // first odd
			mid = i
		}
		c := byte(i + 'A')
		for j, n := 0, v/2; j < n; j++ {
			ans = append(ans, c)
		}
	}

	n := len(ans)
	if mid != -1 {
		c := byte(mid + 'A')
		ans = append(ans, c)
	}

	for i := n - 1; i >= 0; i-- {
		ans = append(ans, ans[i])
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanIntLn(br, &n)

	freq := make([]int, 26)
	for i := 0; i < n; i++ {
		c, _ := br.ReadByte()
		freq[c-'A']++
	}

	ans := solve(freq)
	bw.Write(ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
