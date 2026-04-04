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
			`push 1
back
exit`,
			`ok
1
bye`,
			true,
		},
		{
			"2",
			`size
push 1
size
push 2
size
push 3
size
exit`,
			`0
ok
1
ok
2
ok
3
bye`,
			true,
		},
		{
			"3",
			`push 3
push 14
size
clear
push 1
back
push 2
back
pop
size
pop
size
exit`,
			`ok
ok
2
ok
ok
1
ok
2
2
1
1
0
bye`,
			true,
		},
		{
			"x4",
			`back
pop
exit`,
			`error
error
bye`,
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
