package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKOr([]int{7, 12, 9, 8, 9, 15}, 4))
}

func findKOr(nums []int, k int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		cnt := 0
		for _, ch := range nums {
			if (ch>>i)&1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			ans = ans | (1 << i)
		}
	}
	return ans
}
