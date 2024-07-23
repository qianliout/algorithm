package main

func main() {

}

func maxAscendingSum(nums []int) int {
	n := len(nums)
	ans := nums[0]

	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = nums[i]
		if i-1 >= 0 && nums[i-1] < nums[i] {
			f[i] = f[i-1] + nums[i]
		}

		ans = max(ans, f[i])
	}
	return ans
}
