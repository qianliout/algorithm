package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(3, 2, 3, 5))
}

func nthUglyNumber(n int, a int, b int, c int) int {
	ab := lcm(a, b)
	ac := lcm(a, c)
	bc := lcm(b, c)
	abc := lcm(ab, c)

	var check func(mid int) bool
	check = func(mid int) bool {
		cnt := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		return cnt >= n
	}

	mi := min(a, b, c)
	mx := mi*n + 1
	le, ri := mi, mx
	for le < ri {
		mid := le + (ri-le)/2
		if le >= mi && le < mx && check(mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func lcm(a, b int) int {
	return a * b / gcb(a, b)
}

func gcb(a, b int) int {
	if b == 0 {
		return a
	}
	return gcb(b, a%b)
}
