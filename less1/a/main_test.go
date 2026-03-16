package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run_solve(t *testing.T) {
	test_run(t, solve)
}

func test_run(t *testing.T, solve solveFunc) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		debug   bool
	}{
		{
			"1 пересекаются в точке",
			args{strings.NewReader(`0 7
12 5`)},
			`25`,
			true,
		},
		{
			"не пересекаются",
			args{strings.NewReader(`1 2
-10 3`)},
			`12`,
			true,
		},
		{
			"совпадают",
			args{strings.NewReader(`0 3 0 3`)},
			`7`,
			true,
		},
		{
			"один вложен в другой",
			args{strings.NewReader(`0 2 0 3`)},
			`7`,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debug = v }(debug)
			debug = tt.debug

			out := &bytes.Buffer{}
			run(tt.args.in, out, solve)
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
