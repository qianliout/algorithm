package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(maxDiff(123456))
	fmt.Println(maxDiff(555))
}

// 两次选择，可以选不同的数
func maxDiff(num int) int {
	ss := strconv.Itoa(num)

	mi, mx := num, num
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			res := chang(ss, i, j)
			n, _ := strconv.Atoi(res)
			if res[0] != '0' {
				mi = min(mi, n)
				mx = max(mx, n)
			}
		}
	}

	return mx - mi
}

func chang(ss string, a, b int) string {
	nums := []byte(ss)
	for i, ch := range nums {
		if ch == byte(a+'0') {
			nums[i] = byte(b + '0')
		}
	}
	return string(nums)
}
