package main

func main() {
	nums := []int{-1}
	rotate(nums, 2)
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // 这一步最容易出错
	ratee1(nums, 0, n-1)
	ratee1(nums, 0, k-1)
	ratee1(nums, k, n-1)
}

func ratee1(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
