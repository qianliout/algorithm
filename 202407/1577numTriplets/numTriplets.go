package main

import (
	"sort"
)

func main() {

}

func numTriplets(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	return help(nums1, nums2) + help(nums2, nums1)
}

func help(nums1 []int, nums2 []int) int {
	cnt := make(map[int]int)
	n2 := len(nums2)

	for j := 0; j < n2; j++ {
		for i := j + 1; i < n2; i++ {
			cnt[nums2[i]*nums2[j]]++
		}
	}
	ans := 0
	for _, ch := range nums1 {
		ans += cnt[ch*ch]
	}
	return ans
}
