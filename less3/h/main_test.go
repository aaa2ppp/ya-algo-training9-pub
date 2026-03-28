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
			`5
5 1 1 1 1`,
			`1 1 2`,
			true,
		},
		{
			"2",
			`4
1 2 3 4`,
			`1 2 4`,
			true,
		},
		// {
		// 	"3",
		// 	``,
		// 	``,
		// 	true,
		// },
		// {
		// 	"4",
		// 	``,
		// 	``,
		// 	true,
		// },
		{
			"5",
			`10
1000 1199 200 200 200 200 200 200 200 200`,
			`0 1 6`,
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
