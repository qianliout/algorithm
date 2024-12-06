package main

func main() {

}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	// 表示以nums[i]结尾的最长递增子序列的长度
	f := make([]int, n+1)
	for i := range f {
		f[i] = 1
	}
	// 表示最长递增子序列的长度等 f[i]时的个数
	cnt := make([]int, n+1)
	mx := 0
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if f[j]+1 > f[i] {
					f[i] = f[j] + 1
					cnt[i] = cnt[j]
				} else if f[j]+1 == f[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		mx = max(mx, f[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		if f[i] == mx {
			ans += cnt[i]
		}
	}
	return ans
}
