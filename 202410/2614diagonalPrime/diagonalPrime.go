package main

func main() {

}

func diagonalPrime(nums [][]int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if isPrime(nums[i][i]) {
			ans = max(ans, nums[i][i])
		}
	}
	for i := 0; i < n; i++ {
		if isPrime(nums[i][n-i-1]) {
			ans = max(ans, nums[i][n-i-1])
		}
	}
	return ans
}

func isPrime(n int) bool {
	// 这一步容易出错
	if n < 2 {
		return false
	}
	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}
