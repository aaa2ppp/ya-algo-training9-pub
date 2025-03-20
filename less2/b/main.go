package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) []bool

func solve(nums []int) []bool {
	set := make(map[int]bool, len(nums))
	ans := make([]bool, 0, len(nums))

	for _, n := range nums {
		ans = append(ans, set[n])
		set[n] = true
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var nums []int
	nums, _ = ScanIntsLn(br, nums)

	ans := solve(nums)
	for _, v := range ans {
		if v {
			bw.WriteString("YES\n")
		} else {
			bw.WriteString("NO\n")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
