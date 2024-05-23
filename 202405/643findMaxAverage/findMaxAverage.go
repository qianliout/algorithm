package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{-6662, 5432, -8558, -8935, 8731, -3083, 4115, 9931, -4006, -3284, -3024, 1714, -2825, -2374, -2750, -959, 6516, 9356, 8040, -2169, -9490, -3068, 6299, 7823, -9767, 5751, -7897, 6680, -1293, -3486, -6785, 6337, -9158, -4183, 6240, -2846, -2588, -5458, -9576, -1501, -908, -5477, 7596, -8863, -4088, 7922, 8231, -4928, 7636, -3994, -243, -1327, 8425, -3468, -4218, -364,
		4257, 5690, 1035, 6217, 8880, 4127, -6299, -1831, 2854, -4498, -6983, -677, 2216, -1938, 3348, 4099, 3591, 9076, 942, 4571, -4200, 7271, -6920, -1886, 662, 7844, 3658, -6562, -2106, -296, -3280, 8909, -8352, -9413, 3513, 1352, -8825}, 90))
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	win := 0
	le, ri := 0, 0
	var ans float64 = math.MinInt32
	for le <= ri && ri < len(nums) {
		win += nums[ri]
		ri++
		for ri-le > k {
			win -= nums[le]
			le++
		}
		if ri-le == k {
			ans = max(ans, float64(win))
		}
	}
	return ans / float64(k)
}

func findMaxAverage1(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	win := 0
	le, ri := 0, 0
	var ans float64 = math.MinInt32
	for le <= ri && ri < len(nums) {
		win += nums[ri]
		if ri >= k {
			win -= nums[le]
			le++
		}
		if ri >= k-1 {
			ans = max(ans, float64(win)/float64(k))
		}
		ri++
	}
	return ans
}

func findMaxAverage2(nums []int, k int) float64 {
	sum := 0
	ans := float64(math.MinInt32)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i >= k {
			sum -= nums[i-k]
		}
		if i >= k-1 {
			ans = max(ans, float64(sum))
		}
	}
	return ans / float64(k)
}
