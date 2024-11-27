package main

import (
	"math"

	. "outback/algorithm/common/treenode"
)

func main() {

}

// 后续遍历,把节点值的范围向上传
func isValidBST(root *TreeNode) bool {
	inf := math.MaxInt64 / 10
	// 返回最小值，最大值
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return inf, -inf
		}
		v := root.Val
		lmi, lmx := dfs(root.Left)
		rmi, rmx := dfs(root.Right)
		if v <= lmx || v >= rmi {
			// 这是不合法的
			return -inf, inf
		}

		return min(lmi, v), max(rmx, v)
	}
	_, r := dfs(root)
	return r != inf
}
