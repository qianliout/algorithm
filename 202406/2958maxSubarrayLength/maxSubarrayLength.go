package main

func main() {

}

func maxSubarrayLength(nums []int, k int) int {
	n, cnt := len(nums), make(map[int]int)
	le, ri := 0, 0
	ans := 0
	for le <= ri && ri < n {
		c := nums[ri]
		ri++
		cnt[c]++
		for cnt[c] > k {
			cnt[nums[le]]--
			le++
		}
		ans = max(ans, ri-le)
	}

	return ans
}
