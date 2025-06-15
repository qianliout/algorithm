package main

import (
	"sort"
)

func main() {

}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	i := 0
	for i < n {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
		i++
	}
	return ans
}
