package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(a, b, c, d int) (int, int)

func solve(a, b, c, d int) (int, int) {
	var am, bm, cn, dn int

	// кол-во гарантирующее по крайней мере один элемент
	am = b + 1
	bm = a + 1
	cn = d + 1
	dn = c + 1

	if a == 0 || c == 0 { // нет комплекта
		return bm, dn
	}
	if b == 0 || d == 0 { // нет комплекта
		return am, cn
	}

	// кандидаты
	var mn [][2]int
	mn = append(mn, [2]int{am, cn})
	mn = append(mn, [2]int{bm, dn})
	mn = append(mn, [2]int{max(am, bm), 1})
	mn = append(mn, [2]int{1, max(cn, dn)})

	// выбирем наименьшую сумму
	minSum, minIdx := mn[0][0]+mn[0][1], 0
	for i := range mn {
		s := mn[i][0] + mn[i][1]
		if s < minSum {
			minSum, minIdx = s, i
		}
	}

	return mn[minIdx][0], mn[minIdx][1]
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var a, b, c, d int
	ScanInt(br, &a, &b, &c, &d)

	n, m := solve(a, b, c, d)
	PrintIntLn(bw, n, m)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
