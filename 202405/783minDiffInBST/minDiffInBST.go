package main

import (
	"math"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func minDiffInBST(root *TreeNode) int {
	ans := math.MaxInt

	dfs(root, nil, nil, &ans)
	if ans == math.MaxInt {
		return 0
	}

	return ans
}

func dfs(root *TreeNode, ma, mi *TreeNode, ans *int) {
	if root == nil {
		return
	}

	if ma != nil {
		*ans = min(*ans, ma.Val-root.Val)
	}
	if mi != nil {

		*ans = min(*ans, root.Val-mi.Val)
	}

	dfs(root.Left, root, mi, ans)
	dfs(root.Right, ma, root, ans)
}
