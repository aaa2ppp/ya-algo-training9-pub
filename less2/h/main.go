package main

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type Bank struct {
	accounts map[string]int
}

func NewBank() *Bank {
	return &Bank{make(map[string]int)}
}

func (b *Bank) DEPOSIT(name string, sum int) {
	b.accounts[name] += sum
}

func (b *Bank) WITHDRAW(name string, sum int) {
	b.accounts[name] -= sum
}

func (b *Bank) BALANCE(name string) (int, bool) {
	sum, exists := b.accounts[name]
	return sum, exists
}

func (b *Bank) TRANSFER(name1, name2 string, sum int) {
	b.accounts[name1] -= sum
	b.accounts[name2] += sum
}

func (b *Bank) INCOME(p int) {
	for customer, sum := range b.accounts {
		if sum > 0 {
			sum = sum * (100 + p) / 100
			b.accounts[customer] = sum
		}
	}
}

type solveFunc func()

func solve() {}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	b := NewBank()

	var err error
	var op string
loop:
	for {
		op, err = br.ReadString('\n')
		if err != nil {
			break
		}
		f := strings.Fields(op)
		switch op := f[0]; op {
		case "DEPOSIT":
			name := f[1]
			sum, _ := strconv.Atoi(f[2])
			b.DEPOSIT(name, sum)
		case "WITHDRAW":
			name := f[1]
			sum, _ := strconv.Atoi(f[2])
			b.WITHDRAW(name, sum)
		case "BALANCE":
			name := f[1]
			sum, exists := b.BALANCE(name)
			if exists {
				PrintIntLn(bw, sum)
			} else {
				PrintWordLn(bw, "ERROR")
			}
		case "TRANSFER":
			name1 := f[1]
			name2 := f[2]
			sum, _ := strconv.Atoi(f[3])
			b.TRANSFER(name1, name2, sum)
		case "INCOME":
			p, _ := strconv.Atoi(f[1])
			b.INCOME(p)
		default:
			err = errors.New("unknown operation: " + op)
			break loop
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
