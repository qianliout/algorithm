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
