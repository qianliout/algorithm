package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimizedMaximum(1, []int{1}))

}

func minimizedMaximum(n int, quantities []int) int {
	mx := slices.Max(quantities) + 1
	le, ri := 1, mx
	for le < ri {
		mid := le + (ri-le)/2
		// 左端点
		if mid < mx && check(n, quantities, mid) {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return le
}

func check(n int, q []int, mx int) bool {
	cnt := 0
	for _, ch := range q {
		cnt += (ch + mx - 1) / mx
	}
	return cnt <= n
}
