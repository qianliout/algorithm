package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSeconds([]int{11, 4, 10}))
}

func minimumSeconds(nums []int) int {
	cnt := make(map[int][]int)
	for i, ch := range nums {
		cnt[ch] = append(cnt[ch], i)
	}
	n := len(nums)
	ans := n
	for _, idx := range cnt {
		s, e := idx[0], idx[len(idx)-1]
		mx := n - e + s
		for i := 1; i < len(idx); i++ {
			mx = max(mx, idx[i]-idx[i-1])
		}
		ans = min(ans, mx)
	}
	return ans / 2
}
