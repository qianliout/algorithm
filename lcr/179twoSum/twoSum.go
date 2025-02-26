package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 9, 12, 15}, 18))
}

func twoSum(price []int, target int) []int {
	ans := make([]int, 2)
	n := len(price)
	for i := 0; i < n-1; i++ {
		if find(price[i+1:], target-price[i]) {
			ans[0] = price[i]
			ans[1] = target - price[i]
		}
	}
	return ans
}

func find(nums []int, target int) bool {
	n := len(nums)
	le, ri := 0, n
	for le < ri {
		// 找>=target的左端点
		mid := le + (ri-le)/2
		if mid < n && nums[mid] >= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < n && nums[le] == target {
		return true
	}
	return false
}
