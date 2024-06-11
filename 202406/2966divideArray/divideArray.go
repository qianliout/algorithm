package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 3))
}

func divideArray2(nums []int, k int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	i := 0
	for i < n {

		th := make([]int, 0)
		for j := i; j < i+3; j++ {
			// if j == i && len(ans) > 0 {
			// 	if ans[len(ans)-1][2] == nums[j] {
			// 		return [][]int{}
			// 	}
			// }

			if j > 0 && abs(nums[j]-nums[i]) > k {
				return [][]int{}
			}
			th = append(th, nums[j])
		}
		ans = append(ans, th)
		i = i + 3
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		ans = append(ans, append([]int{}, nums[i:i+3]...))
	}
	return ans
}
