package main

func main() {

}

// [1,3,5,4,7] 会得到错的答案
func findNumberOfLIS1(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	mx := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		mx = max(mx, f[i])
	}
	cnt := 0
	for _, ch := range f {
		if ch == mx {
			cnt++
		}
	}
	return cnt
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	counter := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = 1
		counter[i] = 1
	}
	mx := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					counter[i] = counter[j]
				} else if dp[j]+1 == dp[i] {
					counter[i] += counter[j]
				}
			}
		}
		if dp[i] > mx {
			mx = dp[i]
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if dp[i] == mx {
			ans += counter[i]
		}
	}
	return ans
}
