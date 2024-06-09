package main

import (
	"fmt"
)

func main() {
	fmt.Println(findIndices([]int{2, 0, 9, 2}, 2, 4))
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	mx, mi := 0, 0

	j, n := indexDifference, len(nums)
	for j < n {
		i := j - indexDifference
		if nums[i] > nums[mx] {
			mx = i
		}
		if nums[i] < nums[mi] {
			mi = i
		}

		if mx <= j-indexDifference && abs(nums[j]-nums[mx]) >= valueDifference {
			return []int{mx, j}
		}
		if mi <= j-indexDifference && abs(nums[j]-nums[mi]) >= valueDifference {
			return []int{mi, j}
		}
		j++
	}
	return []int{-1, -1}
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
