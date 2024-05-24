package main

import (
	"fmt"
)

func main() {
	fmt.Println(secondGreaterElement([]int{2, 4, 0, 9, 6}))
}
func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	s, t := make([]int, 0), make([]int, 0)
	for i, ch := range nums {
		// 这里是去更新答案，只有严格大于才算答案
		for len(t) > 0 && nums[t[len(t)-1]] < ch {
			ans[t[len(t)-1]] = ch
			t = t[:len(t)-1]
		}
		l := len(s)
		for l > 0 && nums[s[l-1]] < ch {
			l--
		}
		// 整体搬迁到 t 中
		t = append(t, s[l:]...)
		s = s[:l]

		s = append(s, i)
	}

	return ans
}
