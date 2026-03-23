package main

import (
	"io"
	"os"
	"slices"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([][]string) (_, _ []string)

func solve(studentLangs [][]string) (everyone, leastOne []string) {
	n := len(studentLangs)
	langCount := make(map[string]int)

	for _, langs := range studentLangs {
		for _, lang := range langs {
			langCount[lang]++
		}
	}
	for lang, count := range langCount {
		leastOne = append(leastOne, lang)
		if count == n {
			everyone = append(everyone, lang)
		}
	}
	return
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n int
	ScanInt(br, &n)

	studentLangs := make([][]string, 0, n)
	for i := 0; i < n; i++ {
		var m int
		ScanInt(br, &m)
		langs := make([]string, m)
		ScanWords(br, langs)
		studentLangs = append(studentLangs, langs)
	}

	everyone, leastOne := solve(studentLangs)

	// от нас не требуют определенный порядок вывода, но так проще тестировать
	slices.Sort(everyone)
	slices.Sort(leastOne)

	op := WO{Sep: "\n", End: "\n"}

	PrintIntLn(bw, len(everyone))
	PrintWords(bw, op, everyone)

	PrintIntLn(bw, len(leastOne))
	PrintWords(bw, op, leastOne)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
