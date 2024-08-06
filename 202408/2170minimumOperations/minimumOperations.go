package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumOperations([]int{3, 1, 3, 2, 4, 3}))
	fmt.Println(minimumOperations([]int{1, 2, 2, 2, 2}))
	fmt.Println(minimumOperations([]int{1}))
}

func minimumOperations(nums []int) int {
	nums1, nums2 := make([]int, 0), make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i = i + 2 {
		nums1 = append(nums1, nums[i])
		if i+1 < n {
			nums2 = append(nums2, nums[i+1])
		}
	}
	cnt1 := make(map[int]pair)
	for _, ch := range nums1 {
		pa := cnt1[ch]
		pa.val = ch
		pa.cnt++
		cnt1[ch] = pa
	}
	cnt2 := make(map[int]pair)
	for _, ch := range nums2 {
		pa := cnt2[ch]
		pa.val = ch
		pa.cnt++
		cnt2[ch] = pa
	}
	pairs1 := make([]pair, 2)
	for _, ch := range cnt1 {
		pairs1 = append(pairs1, ch)
	}
	sort.Slice(pairs1, func(i, j int) bool { return pairs1[i].cnt > pairs1[j].cnt })

	pairs2 := make([]pair, 2)
	for _, ch := range cnt2 {
		pairs2 = append(pairs2, ch)
	}
	sort.Slice(pairs2, func(i, j int) bool { return pairs2[i].cnt > pairs2[j].cnt })
	pairs1, pairs2 = pairs2[:2], pairs1[:2] // 不够的数数就是 pair{0,0}

	x1, x2, y1, y2 := pairs1[0], pairs1[1], pairs2[0], pairs2[1]
	if x1.val != y1.val {
		return n - (x1.cnt + y1.cnt)
	}

	return min(n-(x1.cnt+y2.cnt), n-(x2.cnt+y1.cnt))
}

type pair struct {
	val int
	cnt int
}
