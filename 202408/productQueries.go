package main

import (
	"math"
)

func main() {

}

func productQueries(n int, queries [][]int) []int {
	nums := make([]int, 0)
	// 它包含 最少 数目的 2 的幂，且它们的和为 n
	for n > 0 {
		lb := n & -n // lowbit 的计算
		nums = append(nums, lb)
		n ^= lb
	}

	mod := int(math.Pow10(9)) + 7
	ans := make([]int, len(queries))
	for j, ch := range queries {
		s, e := ch[0], ch[1]
		cnt := 1
		for i := s; i <= e; i++ {
			cnt = cnt * nums[i]
			cnt = cnt % mod
		}
		ans[j] = cnt
	}
	return ans
}
