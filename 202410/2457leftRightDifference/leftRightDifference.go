package main

func leftRightDifference(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + nums[i]
	}
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + nums[i]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = abs(right[i] - left[i])
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
