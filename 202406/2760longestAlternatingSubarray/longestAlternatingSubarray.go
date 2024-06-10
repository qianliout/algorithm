package main

func main() {

}

func longestAlternatingSubarray(nums []int, threshold int) int {
	n, i, ans := len(nums), 0, 0
	for i < n {
		if nums[i]%2 != 0 || nums[i] > threshold {
			i++
			continue
		}
		start := i
		i++
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		ans = max(ans, i-start)
	}
	return ans
}
