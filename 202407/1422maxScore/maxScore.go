package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00"))
}

func maxScore(s string) int {
	one := 0
	for _, ch := range s {
		if byte(ch) == '1' {
			one++
		}
	}
	ans := 0
	zero := 0
	for i, ch := range s {
		if byte(ch) == '0' {
			zero++
		} else {
			one--
		}
		// 不能是空串
		if i != len(s)-1 {
			ans = max(zero+one, ans)
		}
	}
	return ans
}
