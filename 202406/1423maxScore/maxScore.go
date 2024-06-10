package main

func main() {

}

// 会timeout
func maxScore1(cardPoints []int, k int) int {
	var dfs func(nums []int, k int) int
	dfs = func(nums []int, k int) int {
		if len(nums) == 0 || k <= 0 {
			return 0
		}
		res1 := nums[0] + dfs(nums[1:], k-1)
		res2 := nums[len(nums)-1] + dfs(nums[:len(nums)-1], k-1)
		return max(res2, res1)
	}
	return dfs(cardPoints, k)
}

// 拿走 k 张，剩下 n−k 张。这里 n 是 cardPoints 的长度。
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	sum := 0

	for _, ch := range cardPoints {
		sum += ch
	}
	le, ri := 0, 0
	wind := 0
	sub := sum
	for le <= ri && ri < n {
		wind += cardPoints[ri]
		ri++
		for ri-le > n-k {
			wind -= cardPoints[le]
			le++
		}
		if ri-le == n-k {
			sub = min(sub, wind)
		}
	}
	return sum - sub
}
