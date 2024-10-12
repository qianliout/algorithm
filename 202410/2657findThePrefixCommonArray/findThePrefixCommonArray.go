package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
}

/*
1 <= A.length == B.length == n <= 50
1 <= A[i], B[i] <= n
题目保证 A 和 B 两个数组都是 n 个元素的排列。
*/
func findThePrefixCommonArray(A []int, B []int) []int {
	setA := 0
	setB := 0
	n := len(A)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		setA = setA | (1 << A[i])
		setB = setB | (1 << B[i])
		ans[i] = bits.OnesCount(uint(setA & setB))
	}
	return ans
}
