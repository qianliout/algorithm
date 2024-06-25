package main

import (
	"fmt"
)

func main() {
	fmt.Println(strWithout3a3b(1, 3))
}

func strWithout3a3b(a int, b int) string {
	ans := make([]byte, 0)
	for a > 0 && b > 0 {
		if a == b {
			ans = append(ans, 'a')
			ans = append(ans, 'b')
			a--
			b--
		} else if a > b {
			ans = append(ans, 'a', 'a')
			ans = append(ans, 'b')
			a = a - 2
			b = b - 1
		} else if a < b {
			ans = append(ans, 'b', 'b')
			ans = append(ans, 'a')
			a = a - 1
			b = b - 2
		}
	}
	// 题目中明确有解
	for i := 0; i < a; i++ {
		ans = append(ans, 'a')
	}
	for i := 0; i < b; i++ {
		ans = append(ans, 'b')
	}
	return string(ans)
}
