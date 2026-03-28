package main

import (
	"io"
	"os"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) string

func solve(s string) string {
	freq := make([]int, 10)
	for _, c := range []byte(s) {
		c -= '0'
		freq[c]++
	}

	sum := 0
	for i, v := range freq {
		sum += i * v
	}

	if mod := sum % 3; mod != 0 {
		// пытаемся найти минимальную цифру с остатком от деления на 3 равным mod
		for i := 1; i < 9; i++ {
			if freq[i] == 0 || i%3 != mod {
				continue
			}
			freq[i]--
			sum -= i
			break
		}
	}

	if mod := sum % 3; mod != 0 {
		// есть только два возможных остатка 1 и 2, и нет ни одной цифры с желаемым остатком.
		// если 1, то нужно найти две цифры с остатком 2 (4 % 3 == 1).
		// если 2, то нужно найти две цифры с остатком 1 (2 % 3 == 2).
		// другими словами, найти две цифры, остаток которых не равен 0.
	attempt2:
		for i := 1; i < 9; i++ {
			if freq[i] == 0 || i%3 == 0 {
				continue
			}
			freq[i]--
			sum -= i
			for j := i; j < 9; j++ {
				if freq[j] == 0 || j%3 == 0 {
					continue
				}
				freq[j]--
				sum -= j
				break attempt2
			}
		}
	}

	if mod := sum % 3; mod != 0 {
		panic("cannot solve!") // проверяем себя (нам гарантируют, что есть ответ)
	}

	var ans strings.Builder
	for i := 9; i >= 0; i-- {
		for j := 0; j < freq[i]; j++ {
			ans.WriteByte(byte(i) + '0')
		}
	}

	return ans.String()
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')
	ans := solve(s)
	PrintWordLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
