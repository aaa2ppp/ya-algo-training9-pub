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

type solveFunc func([]int) int

func solve(a []int) int {
	n := len(a)

	// дистанция до ближайшего магазина, определена только для домов (1)
	d := make([]int, n)

	m := -n
	for i := 0; i < n; i++ {
		switch a[i] {
		case 1:
			d[i] = i - m
		case 2:
			m = i
		}
	}

	m = n * 2
	for i := n - 1; i >= 0; i-- {
		switch a[i] {
		case 1:
			d[i] = min(d[i], m-i)
		case 2:
			m = i
		}
	}

	ans := -1
	for i := range d {
		if a[i] == 1 {
			ans = max(ans, d[i])
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	n := 10
	a := make([]int, n)
	if _, err := ScanInts(br, a); err != nil {
		panic(err)
	}

	ans := solve(a)
	PrintIntLn(bw, ans)
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
func appendInt[T Int](buf []byte, v T) []byte {
	signed := ^T(0) < 0
	if signed {
		return strconv.AppendInt(buf, int64(v), 10)
	} else {
		return strconv.AppendUint(buf, uint64(v), 10)
	}
}
func ScanInts[T Int](br *Reader, a []T) (int, error)    { return scanSlice(br, parseInt, a) }
func PrintIntLn[T Int](bw *Writer, a ...T) (int, error) { return printValsLn(bw, appendInt, a...) }
func parseInt[T Int](token []byte) (T, error)           { return parseIntBase[T](token) }

type WO struct {
	Begin string
	Sep   string
	End   string
}
type writeOpts = WO
type appendValFunc[T any] func([]byte, T) []byte

func printSlice[T any](bw *Writer, op writeOpts, appendVal appendValFunc[T], a []T) (int, error) {
	var err error
	var buf []byte
	_, _ = bw.WriteString(op.Begin)
	for i := 0; i < len(a); i++ {
		if bw.Available() < len(bw.scratch) {
			buf = bw.scratch[:0]
		} else {
			buf = bw.AvailableBuffer()
		}
		if i > 0 {
			buf = append(buf, op.Sep...)
		}
		buf = appendVal(buf, a[i])
		if _, err = bw.Write(buf); err != nil {
			return i, err
		}
	}
	_, err = bw.WriteString(op.End)
	return len(a), err
}

var lineWO = WO{Sep: " ", End: "\n"}

func printValsLn[T any](bw *Writer, appendVal appendValFunc[T], a ...T) (int, error) {
	return printSlice(bw, lineWO, appendVal, a)
}

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

func scanSlice[T any](br *Reader, parse parseFunc[T], a []T) (int, error) {
	var eof bool
	for i := range a {
		if eof {
			return i, io.EOF
		}
		if err := skipSpace(br, false); err != nil {
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
		a[i] = v
	}
	return len(a), nil
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
