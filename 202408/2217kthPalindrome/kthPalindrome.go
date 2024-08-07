package main

import (
	"math"
)

func main() {

}

func kthPalindrome(queries []int, intLength int) []int64 {
	ans := make([]int64, len(queries))

	base := int(math.Pow10((intLength - 1) / 2))
	for i, q := range queries {
		if q > 9*base {
			ans[i] = -1
			continue
		}
		pre := base + q - 1
		x := pre
		// 后半部分
		if intLength&1 == 1 {
			x = x / 10
		}
		for x > 0 {
			pre = pre*10 + x%10
			x = x / 10
		}
		ans[i] = int64(pre)
	}
	return ans
}
