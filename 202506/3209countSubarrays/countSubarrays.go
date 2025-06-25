package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countSubarrays([]int{1, 1, 1}, 1))
	fmt.Println(countSubarrays([]int{85, 14, 26, 17, 86, 94}, 14))
	fmt.Println(countSubarrays([]int{1, 0, 10, 10, 4}, 4))
}

func countSubarrays(nums []int, k int) int64 {
	ans := 0
	for i, ch := range nums {
		j := i - 1
		for j >= 0 && nums[j]&ch != nums[j] {
			nums[j] = nums[j] & ch
			j--
		}
		a := nums[:i+1]
		ans += sort.SearchInts(a, k+1) - sort.SearchInts(a, k)

	}
	return int64(ans)
}
