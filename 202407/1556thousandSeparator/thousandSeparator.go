package main

import (
	"fmt"
)

func main() {
	fmt.Println(thousandSeparator(123456789))
}

func thousandSeparator(n int) string {
	str := fmt.Sprintf("%d", n)
	ans := make([]byte, 0)
	start := 0
	for i := len(str) - 1; i >= 0; i-- {
		ans = append(ans, byte(str[i]))
		start++
		if start == 3 {
			ans = append(ans, '.')
			start = 0
		}
	}
	le, ri := 0, len(ans)-1
	for le < ri {
		ans[le], ans[ri] = ans[ri], ans[le]
		le++
		ri--
	}
	if ans[0] == '.' {
		ans = ans[1:]
	}
	return string(ans)
}
