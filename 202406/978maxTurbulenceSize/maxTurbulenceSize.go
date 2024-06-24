package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
}

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	dp1 := make([]int, n) // 偶数
	dp2 := make([]int, n) // 奇数

	for i := 0; i < n; i++ {
		if i&1 == 0 {
			dp1[i] = 1
		} else {
			dp2[i] = 1
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if i&1 == 0 && arr[i] < arr[i-1] {

		}

	}
	return ans
}
