package main

import (
	"sort"
)

func main() {

}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {

	// 已经是有序的了
	n := len(fruits)
	left := sort.Search(n, func(i int) bool { return fruits[i][0] >= startPos-k })
	right, s := left, 0
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1]
	}
	ans := s
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		s += fruits[right][1]
		// 需要移动窗口了
		for left <= right &&
			(fruits[right][0]-startPos+fruits[right][0]-fruits[left][0]) > k &&
			(startPos-fruits[left][0]+fruits[right][0]-fruits[left][0]) > k {

			s -= fruits[left][1]
			left++
		}
		ans = max(ans, s)
	}
	return ans
}
