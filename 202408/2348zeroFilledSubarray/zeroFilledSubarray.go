package main

import (
	"fmt"
)

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
}

func zeroFilledSubarray(nums []int) int64 {
	pre := 0
	ans := 0
	for _, n := range nums {
		if n == 0 {
			ans += 1
			ans += pre
			pre++
		} else {
			pre = 0
		}
	}
	return int64(ans)
}
