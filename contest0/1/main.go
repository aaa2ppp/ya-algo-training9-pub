package main

import (
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(s string) string

func solve(s string) string {
	s = strings.TrimSpace(s) + "  "
	var t strings.Builder
	for i := 0; i < len(s)-2; {
		var n int
		if s[i+2] == '#' {
			n, _ = strconv.Atoi(s[i : i+2])
			i += 3
		} else {
			n = int(s[i]) - '0'
			i++
		}
		t.WriteByte('a' + byte(n) - 1)
	}
	return t.String()
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	s, _ := ScanString(br, '\n')

	s = solve(s)
	bw.WriteString(s)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
