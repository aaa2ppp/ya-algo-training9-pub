package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) int

func solve(s string) int {
	type cell struct{ x, y int }
	visited := make(map[cell]int)
	visited[cell{0, 0}] = 1
	x, y := 0, 0
	count := 0
	for _, c := range []byte(s) {
		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}
		if visited[cell{x, y}] == 1 {
			count++
		}
		visited[cell{x, y}]++
	}
	return count
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')

	ans := solve(s)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
