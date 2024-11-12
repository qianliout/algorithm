package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive1(nums []int) int {
	sort.Ints(nums)
	ans := 0
	cnt := make(map[int]int)
	for _, ch := range nums {
		cnt[ch] = 1
		if cnt[ch-1] > 0 {
			cnt[ch] = max(cnt[ch-1]+1, cnt[ch])
		}
		ans = max(ans, cnt[ch])
	}
	return ans
}

func longestConsecutive(nums []int) int {
	ans := 0
	cnt := make(map[int]int)
	for _, ch := range nums {
		cnt[ch]++
	}
	for _, ch := range nums {
		if cnt[ch-1] <= 0 {
			cru, res := ch, 0
			for cnt[cru] > 0 {
				cnt[cru]-- // 不加这一句也可以得到正确答案，但是会有很多重复的计算
				cru++
				res++
			}
			ans = max(ans, res)
		}
	}
	return ans
}
