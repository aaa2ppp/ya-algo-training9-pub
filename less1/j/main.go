package main

import (
	"bufio"
	"errors"
	"io"
	"math"
	"os"
	"strconv"
	"unsafe"
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
	if _, err := ScanInt(br, &n); err != nil {
		panic(err)

	}

	var a, b, c, d point
	for i := 0; i < n; i++ {
		_, err := ScanInt(br,
			&a.x, &a.y,
			&b.x, &b.y,
			&c.x, &c.y,
			&d.x, &d.y,
		)
		if err != nil {
			panic(err)
		}
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

// -- inline:github.com/aaa2ppp/contestio --------------------------------------

type (
	Sign interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsig interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Int interface{ Sign | Unsig }
)
type IntError struct {
	Num string
	Err error
}

func (e *IntError) Error() string { return "parseInt: " + strconv.Quote(e.Num) + ": " + e.Err.Error() }
func (e *IntError) Unwrap() error { return e.Err }
func parseIntBase[T Int](token []byte) (T, error) {
	var unsigned = ^T(0) >= 0
	orig := token
	if len(orig) == 0 {
		return 0, &IntError{string(orig), strconv.ErrSyntax}
	}
	if orig[0] == '-' || orig[0] == '+' {
		if unsigned {
			return 0, &IntError{string(orig), strconv.ErrSyntax}
		}
		token = token[1:]
		if len(token) == 0 {
			return 0, &IntError{string(orig), strconv.ErrSyntax}
		}
	}
	var u64 uint64
	for _, digit := range token {
		digit -= '0'
		if digit > 9 {
			return 0, &IntError{string(orig), strconv.ErrSyntax}
		}
		if u64 < math.MaxUint64/10 || (u64 == math.MaxUint64/10 && digit <= math.MaxUint64%10) {
			u64 = u64*10 + uint64(digit)
			continue
		}
		return 0, &IntError{string(orig), strconv.ErrRange}
	}
	if unsigned {
		if u64 > uint64(^T(0)) {
			return 0, &IntError{string(orig), strconv.ErrRange}
		}
		return T(u64), nil
	}
	bits := int(unsafe.Sizeof(T(0))) << 3
	absMin := uint64(1) << (bits - 1)
	if orig[0] == '-' {
		if u64 > absMin {
			return 0, &IntError{string(orig), strconv.ErrRange}
		}
		return -T(u64), nil
	}
	if u64 >= absMin {
		return 0, &IntError{string(orig), strconv.ErrRange}
	}
	return T(u64), nil
}
func ScanInt[T Int](br *Reader, a ...*T) (int, error) { return scanVars(br, parseInt, a...) }
func parseInt[T Int](token []byte) (T, error)         { return parseIntBase[T](token) }

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

type parseFunc[T any] func([]byte) (T, error)

func scanVarsCommon[T any](br *Reader, stopAtEol bool, parse parseFunc[T], a ...*T) (int, error) {
	var eof bool
	for i := range a {
		if eof {
			return i, io.EOF
		}
		if err := skipSpace(br, stopAtEol); err != nil {
			return i, err
		}
		token, err := nextToken(br)
		if err != nil {
			if err != io.EOF {
				return i, err
			}
			if len(token) == 0 {
				return i, io.EOF
			}
			eof = true
		}
		v, err := parse(token)
		if err != nil {
			return i, err
		}
		*a[i] = v
	}
	return len(a), nil
}
func scanVars[T any](br *Reader, parser func([]byte) (T, error), a ...*T) (int, error) {
	return scanVarsCommon(br, false, parser, a...)
}

var ErrTokenTooLong = errors.New("token too long")

func nextToken(br *Reader) ([]byte, error) {
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
			if isSpace(buf[i]) {
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
var spaceTab = [256]bool{
	' ':  true,
	'\t': true,
	'\r': true,
	'\n': true,
}

func isSpace(c byte) bool { return spaceTab[c] }
func skipSpace(br *Reader, stopAtEol bool) error {
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
			if !isSpace(c) {
				_, _ = br.Discard(i)
				return nil
			}
		}
		_, _ = br.Discard(len(buf))
		fast = false
	}
}

// -- /inline:github.com/aaa2ppp/contestio -------------------------------------
