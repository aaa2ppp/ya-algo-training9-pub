package main

import (
	"io"
	"log"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]string) bool

func solve(a []string) bool {
	p := make([][]byte, 10)
	for i := range p {
		p[i] = []byte(a[i])
	}
	k := make([]int, 5)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if p[i][j] != '#' {
				continue
			}
			n, ok := check(p, i, j)
			if n == 0 || n > 4 || !ok {
				if debug {
					log.Println(i, j, "fail")
				}
				return false
			}
			if debug {
				log.Println(i, j, "bingo", n)
			}
			k[n]++
		}
	}

	for i := 1; i < 5; i++ {
		if k[i] != 5-i {
			return false
		}
	}

	return true
}

func check(a [][]byte, i, j int) (int, bool) {
	ni, nj := 1, 1

	for i2 := i + 1; i2 < 10 && a[i2][j] == '#'; i2++ {
		a[i2][j] = 'x'
		ni++
	}

	for j2 := j + 1; j2 < 10 && a[i][j2] == '#'; j2++ {
		a[i][j2] = 'x'
		nj++
	}

	if ni > 1 && nj > 1 {
		return 0, false
	}

	if ni > 4 || nj > 4 {
		return 0, false
	}

	{
		i2, j2 := i-1, j-1
		if i2 >= 0 && j2 >= 0 && a[i2][j2] != '.' {
			return 0, false
		}
	}
	{
		i2, j2 := i-1, j+nj
		if i2 >= 0 && j2 < 10 && a[i2][j2] != '.' {
			return 0, false
		}
	}
	{
		i2, j2 := i+ni, j-1
		if i2 < 10 && j2 >= 0 && a[i2][j2] != '.' {
			return 0, false
		}
	}
	{
		i2, j2 := i+ni, j+nj
		if i2 < 10 && j2 < 10 && a[i2][j2] != '.' {
			return 0, false
		}
	}

	return max(ni, nj), true
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	mx := make([]string, 10)
	ScanWords(br, mx)

	if solve(mx) {
		bw.WriteString("YES\n")
	} else {
		bw.WriteString("NO\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
