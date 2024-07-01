package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1))
	fmt.Println(largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2))
}

func largestValsFromLabels(values []int, labels []int, w int, u int) int {
	n := len(values)
	nums := make([]pair, n)
	for i := range values {
		nums[i] = pair{v: values[i], l: labels[i]}
	}
	// 贪心的做法，求最大的和，那就先用最大的值
	sort.Slice(nums, func(i, j int) bool { return nums[i].v > nums[j].v })
	cnt := map[int]int{}
	ans := 0
	for i := 0; i < n && w > 0; i++ {
		pa := nums[i]
		if cnt[pa.l] < u {
			cnt[pa.l]++
			w--
			ans += pa.v
		}
	}

	return ans
}

type pair struct {
	v, l int
}
