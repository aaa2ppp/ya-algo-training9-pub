package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type goal struct {
	a, b   int
	player string
}

type solveFunc func([]string, []goal) (string, int)

func solve(players []string, goals []goal) (string, int) {
	playerScore := make(map[string]int)
	for _, v := range players {
		playerScore[v] = 0
	}

	a, b := 0, 0
	for _, g := range goals {
		if g.a > a {
			playerScore[g.player] += g.a - a
			a = g.a
		}
		if g.b > b {
			playerScore[g.player] += g.b - b
			b = g.b
		}
	}

	var winner string
	var score int = -1
	for k, v := range playerScore {
		if v > score {
			winner = k
			score = v
		}
	}
	return winner, score
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, m int
	var players []string
	var goals []goal

	ScanInt(br, &n)

	players = Resize(players, n)
	ScanWords(br, players)

	ScanInt(br, &m)
	goals = Grow(goals, m)
	for i := 0; i < m; i++ {
		var f1, f2 string
		ScanWord(br, &f1, &f2)
		parts := strings.Split(f1, ":")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		goals = append(goals, goal{a, b, f2})
	}

	winner, score := solve(players, goals)
	fmt.Fprintln(bw, winner, score)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
