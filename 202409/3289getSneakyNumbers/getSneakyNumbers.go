package main

import "fmt"

func main() {
	fmt.Println(getSneakyNumbers([]int{0, 1, 1, 0}))
}

func getSneakyNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != i && nums[nums[i]] != nums[i] {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	n := len(nums)
	return nums[n-2:]
}
