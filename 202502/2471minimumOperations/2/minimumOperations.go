package main

import (
	"fmt"
	"sort"

	. "outback/algorithm/common/treenode"
)

func main() {
	fmt.Println(help([]int{7, 6, 8, 5}))
}

func minimumOperations(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		lev1 := make([]*TreeNode, 0)
		lev2 := make([]int, 0)
		for _, no := range queue {
			lev2 = append(lev2, no.Val)
			if no.Left != nil {
				lev1 = append(lev1, no.Left)
			}
			if no.Right != nil {
				lev1 = append(lev1, no.Right)
			}
		}
		ans += help(lev2)
		queue = lev1
	}

	return ans
}

// 这是个很容易想到的方法
func help(lev []int) int {
	ans := make([]int, len(lev))
	exit := make(map[int]int)

	for i := range ans {
		ans[i] = lev[i]
		exit[lev[i]] = i
	}

	cnt := 0
	sort.Ints(ans)
	i, n := 0, len(ans)
	for i < n {
		// 说明位置是对的，不用管，直接走到下一个位置
		if ans[i] == lev[i] {
			i++
			continue
		}

		// 位置不对，那么就找到这位置本来的值，进行交换
		next := exit[ans[i]]
		lev[i], lev[next] = lev[next], lev[i]
		// 这一步容易出错，因为值改变后，对应值的位置也会改变
		// 因为没有重复的值，才可以这么做
		exit[lev[i]] = i
		exit[lev[next]] = next

		cnt++
	}
	return cnt
}
