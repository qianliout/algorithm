package main

import (
	"fmt"
	. "outback/algorithm/common/treenode"
	"strings"
)

func main() {

}

func isSubStructure(node1 *TreeNode, node2 *TreeNode) bool {
	var dfs func(node1, node2 *TreeNode) bool
	dfs = func(n1, n2 *TreeNode) bool {
		if n2 == nil {
			return true
		}
		if n1 == nil && n2 != nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}
		return dfs(n1.Left, n2.Left) && dfs(n1.Right, n2.Right)
	}

	if node2 == nil || node1 == nil {
		return false
	}
	if dfs(node1, node2) {
		return true
	}
	return isSubStructure(node1.Left, node2) || isSubStructure(node1.Right, node2)
}

// 得不到正的结果
func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}

	if B == nil || A == nil {
		return false
	}

	a := serializer(A)
	b := serializer(B)
	if strings.Contains(a, b) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func serializer(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	ans := fmt.Sprintf("%d", root.Val)
	left := serializer(root.Left)
	right := serializer(root.Right)
	return fmt.Sprintf("%s,%s,%s", ans, left, right)
}

func decorateRecord2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		lev := make([]int, 0)
		lev2 := make([]*TreeNode, 0)
		for _, no := range queue {
			lev = append(lev, no.Val)
			if no.Left != nil {
				lev2 = append(lev2, no.Left)
			}
			if no.Right != nil {
				lev2 = append(lev2, no.Right)
			}
		}
		ans = append(ans, lev...)
		queue = lev2
	}
	return ans
}

func decorateRecord(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	flag := true // true 的话就是从左到右
	for len(queue) > 0 {
		lev := make([]int, 0)
		lev2 := make([]*TreeNode, 0)
		for _, no := range queue {
			lev = append(lev, no.Val)
			if no.Left != nil {
				lev2 = append(lev2, no.Left)
			}
			if no.Right != nil {
				lev2 = append(lev2, no.Right)
			}
		}
		if !flag {
			le, ri := 0, len(lev)-1
			for le < ri {
				lev[le], lev[ri] = lev[ri], lev[le]
				le++
				ri--
			}
		}
		ans = append(ans, lev)
		flag = !flag
		queue = lev2
	}
	return ans
}
