package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < 64; i++ {
		c := 0
		for _, v := range nums {
			c += (v >> i) & 1
		}
		ans += (c % 3) << i
	}

	return ans
}

func rangeBitwiseAnd(left int, right int) int {
	if left > right {
		left, right = right, left
	}
	move := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		move++
	}
	return left << move
}
