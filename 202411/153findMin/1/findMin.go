package main

func main() {

}

func findMin(nums []int) int {

	// 数组中没有重复值，可以理解成找 小于等于最后一个元素的左边界
	n := len(nums)
	target := nums[n-1]
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n && nums[mid] <= target {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	return nums[le]
}

// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
