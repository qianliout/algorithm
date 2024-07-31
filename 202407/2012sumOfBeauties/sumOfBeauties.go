package main

func main() {

}

func sumOfBeauties(nums []int) int {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = max(pre[i-1], nums[i])
	}
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}
	ans := 0
	for i := 1; i <= n-2; i++ {
		if pre[i-1] < nums[i] && nums[i] < suf[i+1] {
			ans += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans += 1
		}
	}
	return ans
}
