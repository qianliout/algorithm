package main

func main() {}

func removeDuplicates(nums []int) int {
	exit := make(map[int]int)
	n := len(nums)
	for _, n := range nums {
		exit[n]++
	}
	i := 0
	for j := 0; j < n; j++ {
		if exit[nums[j]] == 1 {
			nums[i] = nums[j]
			i++
		}
		exit[nums[j]]--
	}
	return i
}
