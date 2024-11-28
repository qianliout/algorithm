package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	le, ri := 0, n-k+1
	for le < ri {
		mid := le + (ri-le)/2
		// 左端点写法
		if mid >= 0 && mid < n-k+1 && (x-arr[mid] > arr[mid+k]-x) {
			ri = mid
		} else {
			le = mid + 1
		}
	}

	if le < 0 || le >= n-k+1 {
		return []int{}
	}

	return arr[le : le+k]
}
