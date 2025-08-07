package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, 0)
	st := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(st) > 0 && i-st[0] >= k {
			st = st[1:]
		}
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i] {
			st = st[:len(st)-1]
		}
		st = append(st, i)

		if i >= k-1 {
			ans = append(ans, nums[st[0]])
		}
	}
	return ans
}

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。
// 滑动窗口每次只向右移动一位。
// 返回 滑动窗口中的最大值 。
