package main

import (
	"io"
	"os"
	"slices"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type Sale struct {
	Customer string
	Product  string
	Count    int
}

type ProductCount struct {
	Product string
	Count   int
}

type CustomerProducs struct {
	Customer string
	Products []ProductCount
}

type solveFunc func([]Sale) []CustomerProducs

func solve(sales []Sale) []CustomerProducs {
	index := make(map[string]map[string]int)

	for _, s := range sales {
		pc, exists := index[s.Customer]
		if !exists {
			pc = map[string]int{}
			index[s.Customer] = pc
		}
		pc[s.Product] += s.Count
	}

	report := make([]CustomerProducs, 0, len(index))

	for customer, pc := range index {
		producs := make([]ProductCount, 0, len(pc))

		for product, count := range pc {
			producs = append(producs, ProductCount{product, count})
		}

		slices.SortFunc(producs, func(a, b ProductCount) int {
			return strings.Compare(a.Product, b.Product)
		})

		report = append(report, CustomerProducs{customer, producs})
	}

	slices.SortFunc(report, func(a, b CustomerProducs) int {
		return strings.Compare(a.Customer, b.Customer)
	})

	return report
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var sales []Sale

	var s Sale
	for {
		if _, err := ScanAnyLn(br, &s.Customer, &s.Product, &s.Count); err != nil {
			break
		}
		sales = append(sales, s)
	}

	report := solve(sales)

	for _, r := range report {
		PrintWord(bw, WO{End: ":\n"}, r.Customer)
		for _, p := range r.Products {
			PrintWord(bw, WO{End: " "}, p.Product)
			PrintIntLn(bw, p.Count)
		}
	}
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
