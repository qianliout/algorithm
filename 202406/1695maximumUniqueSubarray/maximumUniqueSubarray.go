package main

func main() {

}

func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	ans, le, ri := 0, 0, 0
	cnt := make(map[int]int)
	for le <= ri && ri < n {
		c := nums[ri]
		cnt[c]++
		ri++
		// 这样做检测会超时
		// for !check(cnt){
		for cnt[c] > 1 {
			cnt[nums[le]]--
			le++
		}
		ans = max(ans, sum[ri]-sum[le])
	}
	return ans
}

func check(cnt map[int]int) bool {
	for _, v := range cnt {
		if v > 1 {
			return false
		}
	}
	return true
}
