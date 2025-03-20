package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(g, s string) int

func solve(g, s string) int {
	gf := make([]int, 64)
	sf := make([]int, 64)
	var n int

	for _, c := range []byte(g) {
		gf[c-'@']++
	}
	for _, v := range gf {
		if v == 0 {
			n++
		}
	}

	sfAdd := func(c byte) {
		c -= '@'
		if sf[c] == gf[c] {
			n--
		}
		sf[c]++
		if sf[c] == gf[c] {
			n++
		}
	}
	sfDel := func(c byte) {
		c -= '@'
		if sf[c] == gf[c] {
			n--
		}
		sf[c]--
		if sf[c] == gf[c] {
			n++
		}
	}

	var ans int

	for i := 0; i < len(g); i++ {
		sfAdd(s[i])
	}
	if n == 64 {
		ans++
	}

	for l, r := 0, len(g); r < len(s); l, r = l+1, r+1 {
		sfDel(s[l])
		sfAdd(s[r])
		if n == 64 {
			ans++
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var m, n int
	ScanIntLn(br, &m, &n)

	g, _ := br.ReadString('\n')
	s, _ := br.ReadString('\n')

	ans := solve(g, s)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
