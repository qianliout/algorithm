package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	le, ri := 0, n-k
	for le < ri {
		mid := le + (ri-le)/2
		// 对于中间位置 mid，我们比较 arr[mid] 和 arr[mid+k] 与 x 的距离
		// 如果 |arr[mid] - x| <= |arr[mid+k] - x|，则最优解在左半部分
		// 否则，最优解在右半部分
		//  不可以使用 abs，不知道原因
		// if mid > 0 && mid+k < n && (abs(x-arr[mid]) <= abs(arr[mid+k]-x)) {
		if mid > 0 && mid+k < n && x-arr[mid] <= arr[mid+k]-x {
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
整数 a 比整数 b 更接近 x 需要满足：
|a - x| < |b - x| 或者
|a - x| == |b - x| 且 a < b
*/
