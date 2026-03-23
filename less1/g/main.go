package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(a []int) int {
	n := len(a)

	// дистанция до ближайшего магазина, определена только для домов (1)
	d := make([]int, n)

	m := -n
	for i := 0; i < n; i++ {
		switch a[i] {
		case 1:
			d[i] = i - m
		case 2:
			m = i
		}
	}

	m = n * 2
	for i := n - 1; i >= 0; i-- {
		switch a[i] {
		case 1:
			d[i] = min(d[i], m-i)
		case 2:
			m = i
		}
	}

	ans := -1
	for i := range d {
		if a[i] == 1 {
			ans = max(ans, d[i])
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	n := 10
	a := make([]int, n)
	ScanInts(br, a)

	ans := solve(a)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
