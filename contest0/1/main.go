package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unsafe"
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

	s, err := br.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	s = solve(s)
	bw.WriteString(s)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}

// -- inline:github.com/aaa2ppp/contestio --------------------------------------

const defaultBufSize = 4096

type Reader = bufio.Reader

func NewReaderSize(r io.Reader, size int) *Reader {
	return bufio.NewReaderSize(r, size)
}
func NewReader(r io.Reader) *Reader {
	return NewReaderSize(r, defaultBufSize)
}

type Writer struct {
	*bufio.Writer
	scratch [64 - unsafe.Sizeof(uintptr(0))]byte
}

func NewWriterSize(w io.Writer, size int) *Writer {
	return &Writer{Writer: bufio.NewWriterSize(w, size)}
}
func NewWriter(w io.Writer) *Writer {
	return NewWriterSize(w, defaultBufSize)
}

// -- /inline:github.com/aaa2ppp/contestio -------------------------------------
