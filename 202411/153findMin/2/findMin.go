package main

func main() {

}

func findMin(nums []int) int {
	n := len(nums)
	le, ri := 0, n
	for le < ri {
		mid := le + (ri-le)/2
		if mid >= 0 && mid < n && nums[mid] < nums[0] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	if le < 0 || le >= n {
		return nums[0]
	}
	return nums[le]
}

// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
