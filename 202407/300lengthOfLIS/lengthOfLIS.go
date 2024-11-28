package main

func main() {

}

func lengthOfLIS(nums []int) int {
	ans := 0
	n := len(nums)
	f := make([]int, n)
	for i, ch := range nums {
		f[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < ch {
				f[i] = max(f[i], f[j]+1)
			}
		}
		ans = max(ans, f[i])
	}
	return ans
}
