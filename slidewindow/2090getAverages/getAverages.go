package main

func main() {

}

func getAverages2(nums []int, k int) []int {
	le, ri, n := 0, 0, len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	win := 0
	for le <= ri && ri < n {
		win += nums[ri]
		ri++
		if ri-le >= 2*k+1 {
			ans[ri-1-k] = win / (2*k + 1)
		}
		if ri-le >= 2*k+1 {
			win -= nums[le]
			le++
		}
	}
	return ans
}

func getAverages(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	pre := make([]int, n+1)
	for i, ch := range nums {
		pre[i+1] = pre[i] + ch
	}
	for i := 2*k + 1; i <= n; i++ {
		sm := pre[i] - pre[i-2*k-1]
		ans[i-1-k] = sm / (2*k + 1)
	}
	return ans
}
