package main

import (
	"io"
	"log"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(a []int) int {
	type item struct {
		val   int
		count int
	}

	var ans int

	var st Stack[item]
	for i, c := range a {
		if debug {
			log.Printf("%d: %d %v", i, c, st)
		}
		if st.Empty() {
			st.Push(item{c, 1})
			continue
		}
		t := st.Top()
		if t.val == c {
			st.Push(item{c, t.count + 1})
			continue
		}
		if t.count >= 3 {
			ans += t.count
			for n := t.count; n > 0; n-- {
				st.Pop()
			}
		}
		if st.Empty() {
			st.Push(item{c, 1})
			continue
		}
		t = st.Top()
		if t.val == c {
			st.Push(item{c, t.count + 1})
			continue
		}
		st.Push(item{c, 1})
	}

	if !st.Empty() {
		t := st.Top()
		if t.count >= 3 {
			ans += t.count
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	a := make([]int, n)
	ScanInts(br, a)

	ans := solve(a)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
