package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func([]int) int

func solve(a []int) int {
	n := len(a)
	_ = n
	return 0
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var cmd string
	var val int
	var stack Stack[int]
	for {
		ScanWord(br, &cmd)
		switch cmd {
		case "push":
			ScanInt(br, &val)
			stack.Push(val)
			PrintWordLn(bw, "ok")
		case "pop":
			if stack.Empty() {
				PrintWordLn(bw, "error")
				break
			}
			PrintIntLn(bw, stack.Pop())
		case "back":
			if stack.Empty() {
				PrintWordLn(bw, "error")
				break
			}
			PrintIntLn(bw, stack.Top())
		case "size":
			PrintIntLn(bw, stack.Len())
		case "clear":
			stack.Reset()
			PrintWordLn(bw, "ok")
		case "exit":
			PrintWordLn(bw, "bye")
			return
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
