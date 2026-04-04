package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(string) (int, error)

func solve(expr string) (int, error) {
	prior := [128]int{
		'(': 1,
		'+': 2,
		'-': 2,
		'*': 4,
	}
	var polish []string
	var st Stack[byte]

	wantExpr := true
	i := skipSpace(expr)
	for i < len(expr) {
		n, err := func() (int, error) {
			c := expr[i]

			if wantExpr {
				if c == '(' {
					st.Push('(')
					wantExpr = true
					return 1, nil
				}
				v, err := scanInt(expr[i:])
				if err != nil {
					return 0, err
				}
				polish = append(polish, v)

				wantExpr = false
				return len(v), nil
			}

			if c == ')' {
				for !st.Empty() && st.Top() != '(' {
					op := st.Pop()
					polish = append(polish, string(op))
				}
				if st.Empty() {
					return 0, errors.New("unexpected closing parenthesis")
				}
				st.Pop()

				wantExpr = false
				return 1, nil
			}

			p := prior[c]
			if p == 0 {
				return 0, errors.New("operator was expected")
			}

			for !st.Empty() && prior[st.Top()] >= p {
				op := st.Pop()
				polish = append(polish, string(op))
			}
			st.Push(c)

			wantExpr = true
			return 1, nil
		}()

		if err != nil {
			return 0, fmt.Errorf("stop at %d: %w", i, err)
		}

		i += n
		i += skipSpace(expr[i:])
	}

	for !st.Empty() {
		op := st.Pop()
		if op == '(' {
			return 0, errors.New("unclosed parenthesis was found")
		}
		polish = append(polish, string(op))
	}

	if debug {
		log.Printf("%q -> %v", expr, polish)
	}

	return evalPolish(polish)
}

func skipSpace(expr string) int {
	for i, c := range []byte(expr) {
		if !unicode.IsSpace(rune(c)) {
			return i
		}
	}
	return len(expr)
}

func scanInt(expr string) (string, error) {
	if len(expr) == 0 {
		return "", io.EOF
	}
	i := 0
	if expr[0] == '+' || expr[0] == '-' {
		i++
	}
	d := 0
	for i < len(expr) && unicode.IsDigit(rune(expr[i])) {
		i++
		d++
	}
	if d == 0 {
		return "", errors.New("number must contain at least one digit")
	}
	return expr[:i], nil
}

func evalPolish(expr []string) (int, error) {
	var st Stack[int]
	for _, token := range expr {
		switch token {
		case "+":
			if st.Len() < 2 {
				return 0, fmt.Errorf("%q: few arguments", token)
			}
			b, a := st.Pop(), st.Pop()
			st.Push(a + b)
		case "-":
			if st.Len() < 2 {
				return 0, fmt.Errorf("%q: few arguments", token)
			}
			b, a := st.Pop(), st.Pop()
			st.Push(a - b)
		case "*":
			if st.Len() < 2 {
				return 0, fmt.Errorf("%q: few arguments", token)
			}
			b, a := st.Pop(), st.Pop()
			st.Push(a * b)
		default:
			v, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			st.Push(v)
		}
	}
	return st.Pop(), nil
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	expr, _ := ScanString(br, '\n')

	ans, err := solve(expr)
	if err != nil {
		if debug {
			log.Println(err)
		}
		bw.WriteString("WRONG\n")
		return
	}
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
