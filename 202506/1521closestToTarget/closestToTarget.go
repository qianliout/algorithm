package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(closestToTarget([]int{1, 2, 4, 8, 16}, 0))
}

func closestToTarget(nums []int, target int) int {
	inf := int(math.Pow10(9))
	ans := inf
	for i, ch := range nums {
		ans = min(ans, abs(ch-target))
		j := i - 1
		for j >= 0 && nums[j]&ch != nums[j] {
			nums[j] = nums[j] & ch
			ans = min(ans, abs(target-nums[j]))
			j--
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func help(arr []int, l, r int) int {
	if r < l {
		return -100000
	}
	ans := arr[l]
	for i := l + 1; i <= r; i++ {
		ans = ans & arr[i]
	}
	return ans
}
