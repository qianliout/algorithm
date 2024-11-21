package main

import (
	"fmt"
)

func main() {
	fmt.Println(minImpossibleOR([]int{2, 1}))
	fmt.Println(minImpossibleOR([]int{5, 3, 2}))
}

func minImpossibleOR(nums []int) int {
	exit := make(map[int]int)
	for _, ch := range nums {
		exit[ch] = 1
	}
	i := 0
	for i <= 63 {
		a := 1 << i
		if exit[a] <= 0 {
			return a
		}
		i++
	}

	return -1
}

// 由于 or 运算只能把二进制中的 0 变成 1（而无法反向操作），因此如果 2k 是可表达的，那么 nums 里一定存在 2k
