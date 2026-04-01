package main

func main() {}

func searchRange(nums []int, target int) []int {
	ans := []int{searchLeft(nums, target), searchRight(nums, target)}
	return ans
}

func searchLeft(nums []int, target int) int {
	l, r := 0, len(nums)
	n := len(nums)
	for l < r {
		mid := l + (r-l)/2
		// 找大于等于target的最小边界
		if mid >= 0 && mid < n && nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l < 0 || l >= len(nums) || nums[l] != target {
		return -1
	}

	return l
}
func searchRight(nums []int, target int) int {
	l, r := 0, len(nums)
	n := len(nums)
	for l < r {
		mid := l + (r-l+1)/2
		// 找大于等于target的最小边界
		if mid >= 0 && mid < n && nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l < 0 || l >= n || nums[l] != target {
		return -1
	}

	return l
}
