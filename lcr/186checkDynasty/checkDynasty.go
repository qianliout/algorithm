package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkDynasty([]int{0, 6, 9, 0, 7}))
}

func checkDynasty(places []int) bool {
	mx, mi := 0, 14
	cnt := make(map[int]int)
	for _, ch := range places {
		if ch == 0 {
			continue
		}
		// 有重复的牌
		if cnt[ch] > 0 {
			return false
		}
		cnt[ch]++
		mx = max(mx, ch)
		mi = min(mi, ch)
	}
	return mx-mi < 5
}

// 这个题目的关键就是可以排序
// 但是直接排序是错的 [0, 6, 9, 0, 7],因为这里的未知朝代可以插入到7和9之间
func checkDynasty2(places []int) bool {
	n := len(places)
	for end := n - 1; end >= 0; end-- {
		if places[end] > 0 && !check(places, end) {
			return false
		}
	}
	return true
}

func check(nums []int, end int) bool {
	pre := nums[end]
	for i := end - 1; i >= 0; i-- {
		if nums[i] == 0 {
			pre--
			if pre <= 0 {
				return false
			}
		} else if nums[i] == pre-1 {
			pre--
		} else {
			return false
		}
	}
	return true
}
