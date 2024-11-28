package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}

func findClosestElements2(arr []int, k int, x int) []int {
	n := len(arr)
	le, ri := 0, n-k
	for le < ri {
		mid := le + (ri-le)/2
		// 左端点写法
		if mid >= 0 && mid < n-k && (x-arr[mid] > arr[mid+k]-x) {
			le = mid + 1
		} else {
			ri = mid
		}
	}

	if le < 0 || le >= n-k+1 {
		return []int{}
	}

	return arr[le : le+k]
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	le, ri := 0, n-k
	for le < ri {
		mid := le + (ri-le)/2
		// 左端点写法
		// 比较le,ri 两个点的大小，确定是怎么缩小区间
		// 这里le,ri 左右两个点，来确定是需要删除左边还是删除右边
		if mid >= 0 && mid+k < n && (x-arr[mid] <= arr[mid+k]-x) {
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
