package main

func main() {

}

func minOperations(nums []int) int {
	start := nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		c := nums[i]
		ans += max(start+1, c) - c
		start = max(start+1, c)
	}

	return ans
}

func minOperations2(nums []int) int {
	start := 0
	ans := 0
	for _, c := range nums {
		ans += max(start+1, c) - c
		start = max(start+1, c)
	}

	return ans
}
