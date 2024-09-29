package main

import "fmt"

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 2, 5}, 3)) // [3,4,-1,-1,-1]
	fmt.Println(resultsArray([]int{3, 2, 3, 2, 3, 2}, 2))    // -1,3,-1,3,-1
}

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] == 1 {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	ans := make([]int, 0)
	for i := k - 1; i < n; i++ {
		if left[i] >= k {
			ans = append(ans, nums[i])
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}
