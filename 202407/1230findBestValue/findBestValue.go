package main

import (
	"fmt"
	"slices"
)

func main() {
	// fmt.Println(findBestValue([]int{4, 9, 3}, 10))
	fmt.Println(findBestValue([]int{2, 3, 5}, 11))
}

// 找各大于等 taget 的左端点，找到之后再试一下  ans-1，就是结果
func findBestValue(arr []int, target int) int {

	mx := slices.Max(arr)
	// 原数组的和都小于等 target 那就直接返回原数组的最大值就行
	if check(arr, mx) <= target {
		return mx
	}

	le, ri := 0, target
	for le < ri {
		mid := le + (ri-le)/2
		// 找左端点
		if le < ri && ri < target+1 && check(arr, mid) >= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	a := check(arr, le)
	b := check(arr, le-1)
	if abs(a-target) < abs(b-target) {
		return le
	}
	return le - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func check(arr []int, mid int) int {
	ans := 0
	for _, ch := range arr {
		ans += min(ch, mid)
	}
	return ans
}
