package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSum([]int{2, 6, 5, 8}, 2))
	fmt.Println(maxSum([]int{4, 5, 4, 7}, 100))
}

/* 选择两个互不相同的下标 i 和 j ，同时 将 nums[i] 更新为 (nums[i] AND nums[j]) 且将 nums[j] 更新为 (nums[i] OR nums[j]) ，OR 表示按位 或 运算，AND 表示按位 与 运算。
本质是把小的数的1放到大的数据中去
*/

func maxSum1(nums []int, k int) int {
	bc := 63
	base := int(math.Pow10(9)) + 7
	bitsCnt := make([]int, bc+1)
	for _, ch := range nums {
		for i := 0; i <= bc; i++ {
			bitsCnt[i] += (ch >> i) & 1
		}
	}

	res := 0
	for k > 0 {
		num := 0
		for i := 0; i <= bc; i++ {
			ch := bitsCnt[i]
			if ch > 0 {
				num |= 1 << i
				bitsCnt[i]--
			}
		}

		res = (res + num*num) % base
		k--
	}

	return res % base
}

func maxSum(nums []int, k int) int {
	bc := 63
	base := int(math.Pow10(9)) + 7
	bitsCnt := make([]int, bc+1)
	for _, ch := range nums {
		for i := 0; i <= bc; i++ {
			bitsCnt[i] += (ch >> i) & 1
		}
	}

	res := 0
	for k > 0 {
		num := 0
		for i, ch := range bitsCnt {
			if ch > 0 {
				num |= 1 << i
				bitsCnt[i]--
			}
		}
		res = (res + num*num) % base
		k--
	}

	return res % base
}
