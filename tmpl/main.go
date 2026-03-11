package main

import (
	"io"
	"log"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) []int

func solve(a []int) []int {
	n := len(a)
	_ = n
	return nil
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	var a []int

	if _, err := ScanIntLn(br, &n); err != nil {
		log.Fatalf("scan n: %v", err)
	}
	if debug {
		log.Printf("n: %d\n", n)
	}

	a = Resize(a, n)
	if i, err := ScanInts(br, a); err != nil {
		log.Printf("scan a[%d]: %v", i, err)
	}
	if debug {
		log.Printf("a: %v", a)
	}

	ans := solve(a)

	PrintIntsLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
