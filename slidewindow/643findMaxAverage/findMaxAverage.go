package main

func main() {

}

func findMaxAverage(nums []int, k int) float64 {
	// 易错点1,这里的初值一定不能是0，因为平均值可能是负数
	ans := -1 << 32

	le, ri, n := 0, 0, len(nums)
	if n < k {
		return 0
	}
	wi := 0
	for le <= n && ri < n {
		wi += nums[ri]
		ri++
		if ri-le == k {
			ans = max(ans, wi)
		}
		// 这里出窗口的条件是：r-le>=k,而不是>k,
		for ri-le >= k {
			wi -= nums[le]
			le++
		}
	}
	return float64(ans) / float64(k)
}
