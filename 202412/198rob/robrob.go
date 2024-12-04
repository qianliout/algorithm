package main

func main() {

}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	f0 := make([]int, n) // not
	f1 := make([]int, n) // yes
	f1[0] = nums[0]
	for i := 1; i < n; i++ {
		f0[i] = max(f0[i-1], f1[i-1])
		f1[i] = max(f0[i-1] + nums[i])
	}
	return max(f0[n-1], f1[n-1])
}
