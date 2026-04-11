package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
)

var debug bool

type solveFunc func([]string) bool

func solve(a []string) bool {
	p := make([][]byte, 10)
	for i := range p {
		p[i] = []byte(a[i])
	}
	k := make([]int, 5)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if p[i][j] != '#' {
				continue
			}
			n, ok := check(p, i, j)
			if n == 0 || n > 4 || !ok {
				if debug {
					log.Println(i, j, "fail")
				}
				return false
			}
			if debug {
				log.Println(i, j, "bingo", n)
			}
			k[n]++
		}
	}

	for i := 1; i < 5; i++ {
		if k[i] != 5-i {
			return false
		}
	}

	return true
}

func check(a [][]byte, i, j int) (int, bool) {
	ni, nj := 1, 1

	for i2 := i + 1; i2 < 10 && a[i2][j] == '#'; i2++ {
		a[i2][j] = 'x'
		ni++
	}

	for j2 := j + 1; j2 < 10 && a[i][j2] == '#'; j2++ {
		a[i][j2] = 'x'
		nj++
	}

	if ni > 1 && nj > 1 {
		return 0, false
	}

	if ni > 4 || nj > 4 {
		return 0, false
	}

	{
		i2, j2 := i-1, j-1
		if i2 >= 0 && j2 >= 0 && a[i2][j2] != '.' {
			return 0, false
		}
	}
	{
		i2, j2 := i-1, j+nj
		if i2 >= 0 && j2 < 10 && a[i2][j2] != '.' {
			return 0, false
		}
	}
	{
		i2, j2 := i+ni, j-1
		if i2 < 10 && j2 >= 0 && a[i2][j2] != '.' {
			return 0, false
		}
	}
	{
		i2, j2 := i+ni, j+nj
		if i2 < 10 && j2 < 10 && a[i2][j2] != '.' {
			return 0, false
		}
	}

	return max(ni, nj), true
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	mx := make([]string, 10)
	ScanWords(br, mx)

	if solve(mx) {
		bw.WriteString("YES\n")
	} else {
		bw.WriteString("NO\n")
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}

// -- inline:github.com/aaa2ppp/contestio --------------------------------------

func _must[T any](v T, err error) (T, error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
	return v, err
}

const defaultBufSize = 4096

type br = bufio.Reader
type bw = bufio.Writer
type Reader struct{ br }

func NewReaderSize(r io.Reader, size int) *Reader {
	return &Reader{*bufio.NewReaderSize(r, size)}
}
func NewReader(r io.Reader) *Reader {
	return NewReaderSize(r, defaultBufSize)
}

type Writer struct {
	bw
	scratch [32]byte
}

func NewWriterSize(w io.Writer, size int) *Writer {
	return &Writer{bw: *bufio.NewWriterSize(w, size)}
}
func NewWriter(w io.Writer) *Writer {
	return NewWriterSize(w, defaultBufSize)
}

type _parseFunc[T any] func([]byte) (T, error)

func _scanSliceCommon[T any](br *Reader, parse _parseFunc[T], a []T) (int, error) {
	for i := range a {
		err := _skipSpace(br, false)
		if err != nil {
			if err == io.EOF {
				if i == 0 {
					return 0, io.EOF
				}
				return i, io.ErrUnexpectedEOF
			}
			return i, err
		}
		token, err := _nextToken(br)
		if err != nil && err != io.EOF {
			return i, err
		}
		v, err := parse(token)
		if err != nil {
			return i, err
		}
		a[i] = v
	}
	return len(a), nil
}
func _scanSlice[T any](br *Reader, parse _parseFunc[T], a []T) (int, error) {
	return _must(_scanSliceCommon(br, parse, a))
}

var ErrTokenTooLong = errors.New("token too long")

func _nextToken(br *Reader) ([]byte, error) {
	var buf []byte
	var err error
	i := 0
	fast := br.Buffered() > 0
	for i < br.Size() {
		if fast {
			buf, _ = br.Peek(br.Buffered())
		} else {
			buf, err = br.Peek(br.Buffered() + 1)
			if err != nil {
				_, _ = br.Discard(len(buf))
				return buf, err
			}
		}
		buf = buf[:br.Buffered()]
		for ; i < len(buf); i++ {
			if _isSpace(buf[i]) {
				_, _ = br.Discard(i)
				return buf[:i], nil
			}
		}
		fast = false
	}
	_, _ = br.Discard(len(buf))
	return buf, ErrTokenTooLong
}

var EOL = errors.New("EOL")
var _spaceTab = [256]bool{
	' ':  true,
	'\t': true,
	'\r': true,
	'\n': true,
}

func _isSpace(c byte) bool { return _spaceTab[c] }
func _skipSpace(br *Reader, stopAtEol bool) error {
	var buf []byte
	var err error
	fast := br.Buffered() > 0
	for {
		if fast {
			buf, _ = br.Peek(br.Buffered())
		} else {
			buf, err = br.Peek(br.Buffered() + 1)
			if err != nil {
				return err
			}
		}
		buf = buf[:br.Buffered()]
		for i, c := range buf {
			if stopAtEol && c == '\n' {
				_, _ = br.Discard(i + 1)
				return EOL
			}
			if !_isSpace(c) {
				_, _ = br.Discard(i)
				return nil
			}
		}
		_, _ = br.Discard(len(buf))
		fast = false
	}
}
func _parseWord[T ~string](token []byte) (T, error)       { return T(token), nil }
func ScanWords[T ~string](br *Reader, a []T) (int, error) { return _scanSlice(br, _parseWord, a) }

// -- /inline:github.com/aaa2ppp/contestio -------------------------------------
