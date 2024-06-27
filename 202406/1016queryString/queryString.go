package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(queryString("10010111100001110010", 10))
	fmt.Println(queryString("110101011011000011011111000000", 15))
}

func queryString(s string, n int) bool {
	for i := n; i >= 1; i-- {
		b := bit2(i)
		if !strings.Contains(s, b) {
			return false
		}
	}

	return true
}

func bit2(n int) string {
	ans := make([]string, 0)
	for n > 0 {
		ans = append(ans, fmt.Sprintf("%d", n%2))
		n = n / 2
	}
	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}

	return strings.Join(ans, "")
}
