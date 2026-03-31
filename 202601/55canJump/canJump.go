package main

func main() {}

func canJump(nums []int) bool {
	n := len(nums)
	mx := 0
	for i := 0; i < n; i++ {
		if mx < i {
			return false
		}
		mx = max(mx, i+nums[i])
	}
	// return mx >= n-1
	return true // 可以直接返回true
}
