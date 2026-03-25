package main

import (
	"io"
	"log"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(k1, m, k2, p2, n2 int) (p1, n1 int)

func solve(k1, m, k2, p2, n2 int) (p1, n1 int) {
	// to 0-indexing
	k1--
	k2--
	p2--
	n2--

	if n2 >= m {
		return -1, -1
	}

	f2 := p2*m + n2 // абсолютный этаж
	if k2 < f2 {
		return -1, -1
	}

	if f2 == 0 { // p2 == 0, n2 == 0
		p1, n1 = -1, -1 // ничего не знаем
		if k1 <= k2 {
			p1, n1 = 0, 0
		}
		if m == 1 { // в доме всего один этаж
			n1 = 0
		}
		if k1 < (k2+1)*m { // номер меньше мин. кол-ва квартир в подъезде
			p1 = 0
		}
		return p1 + 1, n1 + 1
	}

	// f2 * x <= k2 < (f2+1) * x, где х количество квартир на этаже
	// k2/(f2+1) < x <= k2/f2
	x1 := k2/(f2+1) + 1 // нижняя граница
	x2 := k2 / f2       // верхняя граница

	p1, n1 = -2, -2 // пока не нашли ни одного решения

	for x := x1; x <= x2; x++ {
		if k2/x != f2 { // противоречие
			continue
		}

		if k1 == k2 { // однозначное решение
			p1, n1 = p2, n2
			break
		}

		// решаем для x
		f := k1 / x
		p := f / m
		n := f % m

		if debug {
			log.Printf("x:%d p:%d n:%d", x, p, n)
		}

		if p1 == -2 && n1 == -2 { // первое решения
			p1, n1 = p, n
			continue
		}

		if p != p1 { // больше одного решения по подъезду
			p1 = -1
		}
		if n != n1 { // больше одного решения по этажу
			n1 = -1
		}

		if p1 == -1 && n1 == -1 {
			break
		}
	}

	return p1 + 1, n1 + 1 // to 1-indexing
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var k1, m, k2, p2, n2 int
	ScanInt(br, &k1, &m, &k2, &p2, &n2)

	p1, n1 := solve(k1, m, k2, p2, n2)

	PrintIntLn(bw, p1, n1)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
