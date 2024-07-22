package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(closestCost([]int{1, 7}, []int{3, 4}, 10))
	fmt.Println(closestCost([]int{2, 3}, []int{4, 4, 100}, 18))
}

// 没有能获取正确的答案
func closestCost(base []int, topping []int, target int) int {
	// 选第不种基料，选集合 j 的配料
	top2 := append(topping, topping...)
	n := len(base)
	m := len(top2)
	ans := 0

	for i := 0; i < n; i++ {
		cost := base[i]
		for j := 0; j < 1<<m; j++ {

			if bits.OnesCount(uint(j)) > 2 {
				continue
			}

			cost = cal(top2, j)
			if abs(target-cost) < abs(target-ans) || (abs(target-cost) == abs(target-ans) && cost < ans) {
				ans = cost
			}
		}

	}
	return ans
}

func cal(nums []int, k int) int {
	ans := 0
	i := 0
	j := k
	for j > 0 {
		if j&1 == 1 {
			ans += nums[i]
		}
		j = j >> 1
		i++
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
