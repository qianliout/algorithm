package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	i, red, blue := 0, 0, len(nums)-1
	for i < blue {
		if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		} else if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
			i++
		} else {
			i++
		}
	}
}
