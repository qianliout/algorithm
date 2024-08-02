package main

import (
	"fmt"
)

func main() {
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{2, 3}}))
}

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	ret := make([]int, n-1)
	sum := make([]int, n)
	for i := 0; i < n-1; i++ {
		sum[i+1] = sum[i]
		if nums[i]&1 == nums[i+1]&1 {
			ret[i] = 1
			sum[i+1] += 1
		}
	}
	m := len(queries)

	ans := make([]bool, m)
	for i, ch := range queries {
		l, r := ch[0], ch[1]
		if sum[r]-sum[l] > 0 {
			ans[i] = false
		} else {
			ans[i] = true
		}
	}
	return ans
}
