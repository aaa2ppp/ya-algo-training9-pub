package main

import (
	"io"
	"os"

	. "github.com/aaa2ppp/contestio"
)

var debug bool

type solveFunc func(int, []int) int

func solve(k int, a []int) int {
	n := len(a)
	const INF = int(1e12)

	ans := 0
	sum := 0    // текущая сумма окна [l,r]
	add := -INF // макс. некратная сумма отрезка, заканчивающегося на l-1 (всегда отрицательна)
	sub := INF  // мин. положительная некратная сумма префикса [l, r]

	for l, r := 0, 0; r < n; r++ {
		sum += a[r]

		if sum < 0 {
			// ищем лучший add среди суффиксов сбрасываемого блока [l, r]
			sum2 := sum
			add2 := -INF
			for ; l <= r; l++ {
				if sum2 > add2 && sum2%k != 0 {
					add2 = sum2
				}
				sum2 -= a[l]
			}
			add = max(add2, add+sum)
			sum = 0
			sub = INF
			continue
		}

		if sum < sub && sum%k != 0 {
			sub = sum
		}

		if sum > ans {
			if sum%k != 0 {
				ans = sum
			} else {
				// убираем префикс sub или добавляем левый отрезок add
				ans = max(ans, sum-sub)
				ans = max(ans, sum+add)
			}
		}
	}
	return ans
}

func run(in io.Reader, out io.Writer, solve solveFunc) {
	br := NewReader(in)
	bw := NewWriter(out)
	defer bw.Flush()

	var n, k int
	ScanInt(br, &n, &k)

	a := make([]int, n)
	ScanInts(br, a)

	ans := solve(k, a)
	PrintIntLn(bw, ans)
}

func main() {
	run(os.Stdin, os.Stdout, solve)
}
