package main

func main() {

}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	sum := make([]int, n+1)
	for i, ch := range nums {
		sum[i+1] = sum[i] + ch
	}
	ans := make([]int, len(nums))
	for i, ch := range nums {
		pre := -sum[i+1] + i*ch
		nex := sum[n] - sum[i] - (n-i-1)*ch
		ans[i] = pre + nex
	}
	return ans
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
