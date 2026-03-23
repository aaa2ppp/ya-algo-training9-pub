package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type point struct {
	x, y int
}

type solveFunc func(a, b, c, d point) bool

func solve(a, b, c, d point) bool {
	p1 := a
	p2 := b
	p3 := c
	p4 := d

	// мы угадаем, что взяли сторону, а не диоганаль, максимум с двух попыток
	for i := 0; i < 2; i++ {
		// упорядочиваем точки отрезков для сравнения
		if p1.x != p2.x {
			if p1.x > p2.x {
				p1, p2 = p2, p1
			}
			if p3.x > p4.x {
				p3, p4 = p4, p3
			}
		} else {
			if p1.y > p2.y {
				p1, p2 = p2, p1
			}
			if p3.y > p4.y {
				p3, p4 = p4, p3
			}
		}

		if p1.x-p2.x == p3.x-p4.x && p1.y-p2.y == p3.y-p4.y { // параллельны и равны по длине
			return true
		}

		// меняем вторую точку
		p2, p3 = p3, p2
	}

	return false
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	var a, b, c, d point
	for i := 0; i < n; i++ {
		ScanInt(br,
			&a.x, &a.y,
			&b.x, &b.y,
			&c.x, &c.y,
			&d.x, &d.y,
		)
		if solve(a, b, c, d) {
			bw.WriteString("YES\n")
		} else {
			bw.WriteString("NO\n")
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
