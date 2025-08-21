package main

func main() {

}

func minSubArrayLen(target int, nums []int) int {
	wind := make([]int, 0)
	cnt := 0
	n := len(nums)
	ans := n + 1
	for i := 0; i < n; i++ {
		wind = append(wind, i)
		cnt += nums[i]
		for cnt >= target {
			ans = min(ans, i-wind[0]+1)
			cnt -= nums[wind[0]]
			wind = wind[1:]
		}
	}
	if ans == n+1 {
		return 0
	}
	return ans
}

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
