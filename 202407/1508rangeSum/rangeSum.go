package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}
func rangeSum(nums []int, n int, left int, right int) int {
	mod := int(math.Pow10(9)) + 7
	pre1 := make([]int, n+1)
	for i, ch := range nums {
		pre1[i+1] = pre1[i] + ch
	}
	pre2 := make([]int, 0)
	for i := 1; i < len(pre1); i++ {
		for j := i; j < len(pre1); j++ {
			pre2 = append(pre2, pre1[j]-pre1[i-1])
		}
	}
	sort.Ints(pre2)

	sum := make([]int, len(pre2)+1)
	for i, ch := range pre2 {
		sum[i+1] = (sum[i] + ch) % mod
	}
	// fmt.Println(pre1, pre2, sum)
	// 注意下标从1开始
	return sum[right] - sum[left-1]
}
