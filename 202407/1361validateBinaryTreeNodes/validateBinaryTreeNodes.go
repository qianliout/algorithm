package main

import (
	"fmt"
)

func main() {
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
}

/*
如果一棵树是一个二叉树的话 必定除了根节点 每个非空节点的入度都为1
样例还给了某一个节点被两个"父节点"引用的例子 那这个节点的入度就为2了 也就不能构成二叉树
// 特别：一个独立的环 加上 一颗正常的树
*/
func validateBinaryTreeNodes(n int, left []int, right []int) bool {
	in := make([]int, n)
	for i := 0; i < n; i++ {
		if left[i] != -1 {
			in[left[i]]++
		}
		if right[i] != -1 {
			in[right[i]]++
		}
	}
	count0 := 0
	countOther := 0
	rootIndex := 0
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			rootIndex = i
			count0++
		}
		if in[i] > 1 {
			countOther++
		}
	}
	visit := make([]bool, n)
	return count0 == 1 && countOther == 0 && notCircle(rootIndex, left, right, visit) == n
}

// 特别：一个独立的环 加上 一颗正常的树,就得正常判断一下
func notCircle(rootIdx int, left, right []int, visit []bool) int {
	if rootIdx == -1 || visit[rootIdx] {
		return 0
	}
	visit[rootIdx] = true
	le := notCircle(left[rootIdx], left, right, visit)
	ri := notCircle(right[rootIdx], left, right, visit)
	return 1 + le + ri
}
