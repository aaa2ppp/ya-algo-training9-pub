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
			`10 20
heat`,
			"20",
			true,
		},
		{
			"2",
			`20 10
heat`,
			"20",
			true,
		},
		{
			"3",
			`-10 -10
heat`,
			"-10",
			true,
		},
		{
			"4",
			`10 20
freeze`,
			"10",
			true,
		},
		{
			"5",
			`20 10
freeze`,
			"10",
			true,
		},
		{
			"6",
			`-10 -10
freeze`,
			"-10",
			true,
		},
		{
			"7",
			`10 20
auto`,
			"20",
			true,
		},
		{
			"8",
			`20 10
auto`,
			"10",
			true,
		},
		{
			"9",
			`-10 -10
auto`,
			"-10",
			true,
		},
		{
			"10",
			`10 20
fan`,
			"10",
			true,
		},
		{
			"11",
			`20 10
fan`,
			"20",
			true,
		},
		{
			"12",
			`-10 -10
fan`,
			"-10",
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
