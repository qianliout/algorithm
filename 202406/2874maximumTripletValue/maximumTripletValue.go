package main

func main() {

}

// 我们可以在遍历的过程中，维护 nums[i]\textit{nums}[i]nums[i] 的最大值 preMax\textit{preMax}preMax，同时维护 preMax\textit{preMax}preMax 减当前元素的最大值 maxDiff\textit{maxDiff}maxDiff，这就是 kkk 左边 nums[i]−nums[j]\textit{nums}[i] - \textit{nums}[j]nums[i]−nums[j] 的最大值。
func maximumTripletValue1(nums []int) int64 {
	n := len(nums)
	preMax := nums[0]
	difMax := 0
	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]*difMax)
		preMax = max(preMax, nums[i])
		difMax = max(preMax, preMax-nums[i])
	}
	return int64(ans)
}

// 方法一：枚举 j
func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	pre, suf := make([]int, n), make([]int, n) // 前缀最大值和后缀最大值
	pre[0], suf[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		pre[i] = max(pre[i-1], nums[i])
	}
	for i := n - 2; i >= 0; i-- {
		suf[i] = max(suf[i+1], nums[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		ans = max(ans, (pre[i-1]-nums[i])*suf[i+1])
	}
	return int64(ans)
}
