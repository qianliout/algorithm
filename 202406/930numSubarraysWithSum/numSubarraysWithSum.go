package main

func main() {

}

func numSubarraysWithSum(nums []int, goal int) int {
	// 滑动窗口+前缀和（思想）
	// 和为goal的连续子数组的数量。

	// 「和恰好为goal的子数组数量」=「和最多为goal的子数组数量」-「和最多为goal-1的子数组数量」
	return atMost(nums, goal) - atMost(nums, goal-1)

}

// 不大于 goal 的非空子数组

func atMost(nums []int, goal int) int {
	// nums里都是正数
	if goal < 0 {
		return 0
	}
	le, ri := 0, 0
	ans := 0
	sum := 0
	for le <= ri && ri < len(nums) {
		sum += nums[ri]

		ri++
		for le <= ri && sum > goal {
			sum -= nums[le]
			le++
		}
		ans += ri - le
	}

	return ans
}
