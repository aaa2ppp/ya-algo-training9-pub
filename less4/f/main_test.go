package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test_run_solve(t *testing.T) {
	test_run(t, solve)
}

func test_run(t *testing.T, solve solveFunc) {
	tests := []struct {
		name    string
		input   string
		wantOut string
		debug   bool
	}{
		{
			"1",
			`1+(2*2 - 3)`,
			`2`,
			true,
		},
		{
			"x1",
			`1+2*2 - 3`,
			`2`,
			true,
		},
		{
			"2",
			`1+a+1`,
			`WRONG`,
			true,
		},
		{
			"x2",
			`1+1+1`,
			`3`,
			true,
		},
		{
			"x2",
			`1 + a + 1`,
			`WRONG`,
			true,
		},
		{
			"3",
			`1 1 + 2`,
			`WRONG`,
			true,
		},
		{
			"x4",
			`-1--1`,
			`0`,
			true,
		},
		{
			"x5",
			`+1++1`,
			`2`,
			true,
		},
		{
			"x6",
			`2**2`,
			`WRONG`,
			true,
		},
		{
			"7",
			"(1+(7+8)-(3-4*5)*(2*4)+(9)-(16728)*(123*9+2)+((2)))+((((((0-19283))))))",
			`-18570472`,
			true,
		},
		{
			"x8",
			"1+",
			"WRONG",
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debug = v }(debug)
			debug = tt.debug

			in := strings.NewReader(tt.input)
			out := &bytes.Buffer{}
			run(in, out, solve)
			if gotOut := out.String(); trimLines(gotOut) != trimLines(tt.wantOut) {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func trimLines(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t\r\n")
	}
	for n := len(lines); n > 0 && lines[n-1] == ""; n-- {
		lines = lines[:n-1]
	}
	return strings.Join(lines, "\n")
}
