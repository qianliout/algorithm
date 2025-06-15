package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	st := make([]int, 0)
	for i := 0; i < 2*n; i++ {
		c := nums[i%n]
		for len(st) > 0 && c > nums[(st[len(st)-1])%n] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			if ans[last%n] == -1 {
				ans[last%n] = i
			}
		}
		st = append(st, i)
	}
	for i := range ans {
		if ans[i] != -1 {
			ans[i] = nums[ans[i]%n]
		}
	}
	return ans
}
