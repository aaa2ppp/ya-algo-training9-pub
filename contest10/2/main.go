package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(n, k int, input []string) []byte

func solve(n, k int, input []string) []byte {
	win := make([][]byte, n)
	cur := 0
	var cb []byte

	for _, v := range input {
		switch v {
		case "Backspace":
			if len(win[cur]) > 0 {
				win[cur] = win[cur][:len(win[cur])-1]
			}
		case "Copy":
			begin := max(0, len(win[cur])-k)
			cb = append(cb[:0], win[cur][begin:]...)
		case "Paste":
			win[cur] = append(win[cur], cb...)
		case "Next":
			cur = (cur + 1) % n
		default:
			win[cur] = append(win[cur], v...)
		}
	}

	begin := max(0, len(win[cur])-k)
	return win[cur][begin:]
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, m, k int
	ScanInt(br, &n, &m, &k)

	input := make([]string, m)
	ScanWords(br, input)

	ans := solve(n, k, input)
	if len(ans) == 0 {
		bw.WriteString("Empty")
	} else {
		bw.Write(ans)
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
