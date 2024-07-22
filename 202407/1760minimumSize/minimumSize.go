package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumSize([]int{9}, 2))
	fmt.Println(minimumSize([]int{2, 4, 8, 2}, 4))
}

func minimumSize(nums []int, maxOperations int) int {
	mx := slices.Max(nums)
	le, ri := 1, mx
	for le < ri {
		// 找最小值，就是找符和条件的左边界
		mid := le + (ri-le)/2
		if le < mx && check(nums, mid, maxOperations) {
			ri = mid
		} else {
			le = mid + 1
		}
	}

	return le
}

// nums中最大是 mx 值，需要操作多少次
func check(nums []int, mx int, op int) bool {
	cnt := 0
	for _, ch := range nums {
		if ch <= mx {
			continue
		}
		// 向上取整，但是分成4段，只需要三次操作，所以要减1
		cnt += (ch+mx-1)/mx - 1
	}
	return cnt <= op
}
