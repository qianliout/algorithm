package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	fmt.Println(maxTurbulenceSize([]int{100}))
}

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	// up 表示以 i 结尾，并且 arr[i-1]<arr[i]的最子数组的长度
	up := make([]int, n)
	// down 表示以 i 结尾，并且 arr[i-1]>arr[i]的最子数组的长度
	down := make([]int, n)
	for i := 0; i < n; i++ {
		up[i] = 1
		down[i] = 1
	}
	ans := 1
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			up[i] = down[i-1] + 1
		}
		if arr[i-1] > arr[i] {
			down[i] = up[i-1] + 1
		}
		ans = max(ans, up[i], down[i])
	}

	return ans
}
