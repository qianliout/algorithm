package main

import (
	"fmt"
)

func main() {
	fmt.Println(canArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	fmt.Println(canArrange([]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3))
}

func canArrange(arr []int, k int) bool {
	cnt := make(map[int]int)
	for _, ch := range arr {
		cnt[(ch%k+k)%k]++
	}
	for key, v := range cnt {
		if key == 0 {
			if v&1 != 0 {
				return false
			}
		} else {
			if cnt[k-key] != v {
				return false
			}
		}
	}

	return true
}
