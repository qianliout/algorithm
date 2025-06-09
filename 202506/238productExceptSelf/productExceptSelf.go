package main

func main() {

}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n+1)
	after := make([]int, n+1)
	pre[0] = 1
	after[n] = 1
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] * nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		after[i] = after[i+1] * nums[i]
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = pre[i] * after[i+1]
	}
	return ans
}
