package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) int

func solve(s string) int {
	freq := make([]int, 26)
	for _, c := range []byte(s) {
		c -= 'a'
		freq[c]++
	}

	maxFreq := 0
	for _, v := range freq {
		maxFreq = max(maxFreq, v)
	}

	most := make([]bool, 26)
	for c := range freq {
		if freq[c] == maxFreq {
			most[c] = true
		}
	}

	seqs := makeMatrix[byte](26, 27)

	for i, c := range []byte(s) {
		c -= 'a'
		if !most[c] {
			continue
		}
		if seqs[c][0] == 0 { // еще не встречали, добавляем по максимуму
			used := make([]bool, 26)
			for j := 0; j < 26 && i+j < len(s); j++ {
				c2 := s[i+j] - 'a'
				// могут встречаться только наиболее частые и только один раз
				if !most[c2] || used[c2] {
					break
				}
				seqs[c][j] = s[i+j]
				used[c2] = true
			}
		} else { // проверяем и обрезаем
			for j := 0; j < 26 && i+j < len(s); j++ {
				if s[i+j] != seqs[c][j] {
					seqs[c][j] = 0
					break
				}
			}
		}
	}

	// выбираем самую длинную
	count := 1
	for c := 0; c < 26; c++ {
		n := 0
		i := 0
		for seqs[c][i] != 0 {
			n++
			i++
		}
		count = max(count, n)
	}

	return count
}

func makeMatrix[T any](n, m int) [][]T {
	buf := make([]T, n*m)
	matrix := make([][]T, n)
	for i, j := 0, 0; i < n; i, j = i+1, j+m {
		matrix[i] = buf[j : j+m]
	}
	return matrix
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
