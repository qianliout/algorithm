package main

import (
	"math"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root, nil, nil)
}

// 后续遍历
// 难点在于两个返回值不一样
/*
每棵子树返回：

	这棵子树的最小节点值。
	这棵子树的最大节点值。
*/
func check2(root *TreeNode) (int, int) {
	if root == nil {
		// 是二叉搜索树
		// 对于最小值，上层是取 min(lmi,node.Val) 这样就能取到node.Val
		// right 值一样的理解
		return math.MaxInt, math.MinInt
	}
	lmi, lmx := check2(root.Left)
	rmi, rmx := check2(root.Right)
	if root.Val <= lmx || root.Val >= rmi {
		// 不再是二叉搜索树，为了让上层的min(lmi,node.Val) 也是一个不合法的值，只能返回无穷小
		return math.MinInt, math.MaxInt
	}
	return min(lmi, root.Val), max(rmx, root.Val)
}

func check(root, mi, mx *TreeNode) bool {
	if root == nil {
		return true
	}
	if mi != nil && root.Val <= mi.Val {
		return false
	}
	if mx != nil && root.Val >= mx.Val {
		return false
	}
	return check(root.Left, mi, root) && check(root.Right, root, mx)
}

func maxSumBST2(root *TreeNode) int {
	var ans int
	var dfs func(root *TreeNode) (int, int, int)
	dfs = func(root *TreeNode) (int, int, int) {
		if root == nil {
			return math.MaxInt, math.MinInt, 0
		}

		lmi, lmx, lsum := dfs(root.Left)
		rmi, rmx, rsum := dfs(root.Right)
		// 说以 root为结点树 不是一个二叉搜索树
		x := root.Val
		if x <= lmx || x >= rmi {
			return math.MinInt, math.MaxInt, 0
		}
		s := lsum + rsum + x
		ans = max(ans, s)
		return min(lmi, x), max(rmx, x), s
	}

	dfs(root)
	return ans
}

func maxSumBST(root *TreeNode) int {
	var ans int
	dfs(root, &ans)
	return ans
}

/*
每棵子树返回：

    这棵子树的最小节点值。
    这棵子树的最大节点值。
    这棵子树的所有节点值之和。
*/

func dfs(root *TreeNode, ans *int) (int, int, int) {
	if root == nil {
		return math.MaxInt, math.MinInt, 0
	}

	lmi, lmx, lsum := dfs(root.Left, ans)
	rmi, rmx, rsum := dfs(root.Right, ans)
	// 说以 root为结点树 不是一个二叉搜索树
	x := root.Val
	if x <= lmx || x >= rmi {
		return math.MinInt, math.MaxInt, 0
	}
	s := lsum + rsum + x
	*ans = max(*ans, s)
	return min(lmi, x), max(rmx, x), s
}
