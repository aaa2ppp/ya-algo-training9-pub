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
			`105`,
			`510`,
			true,
		},
		{
			"2",
			`2222`,
			`222`,
			true,
		},
		{
			"3",
			`000`,
			`000`,
			true,
		},
		{
			"4",
			`54321`,
			`54321`,
			true,
		},
		{
			"x5",
			`00021`,
			`21000`,
			true,
		},
		{
			"x6",
			`2026`,
			`60`,
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
