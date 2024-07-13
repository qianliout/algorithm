package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
}

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
	if right == 0 {
		return 0
	}
	ans := right // 此时可以认为 arr[:right] 都是可以删除的
	left := 0
	for left == 0 || arr[left-1] <= arr[left] {
		for right < n && arr[right] < arr[left] {
			right++
		}
		// 此时可以认为 arr[left:right] 之间 都是可以删除的
		ans = min(ans, right-left-1)
		left++
	}
	return ans
}
