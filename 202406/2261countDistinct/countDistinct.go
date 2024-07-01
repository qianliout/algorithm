package main

import (
	"fmt"
)

func main() {
	fmt.Println(countDistinct([]int{2, 3, 3, 2, 2}, 2, 2))
}

func countDistinct(nums []int, k int, p int) int {
	set := make(map[sub]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		arr, idx, cnt := [200]int{}, 0, 0
		for j := i; j < n; j++ {
			ch := nums[j]
			if ch%p == 0 {
				cnt++
				if cnt > k {
					break
				}
			}
			arr[idx] = ch
			idx++
			set[arr]++
		}
	}
	return len(set)
}

type sub [200]int
