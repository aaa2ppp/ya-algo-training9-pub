package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) (int, int)

func solve(a []int) (int, int) {
	min1, min2 := a[0], a[1]
	if min1 > min2 {
		min1, min2 = min2, min1
	}
	max1, max2 := a[0], a[1]
	if max1 < max2 {
		max1, max2 = max2, max1
	}

	for _, v := range a[2:] {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
		if v > max1 {
			max2 = max1
			max1 = v
		} else if v > max2 {
			max2 = v
		}
	}

	if min1*min2 > max1*max2 {
		return min1, min2
	}
	return max2, max1
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var a []int
	a, _ = ScanIntsLn(br, a)

	p, q := solve(a)
	PrintIntLn(bw, p, q)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
