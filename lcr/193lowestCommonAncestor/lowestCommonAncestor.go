package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 基础情况
	if root == nil || root == p || root == q {
		return root
	}

	// 在左右子树中递归查找
	// 这里的 l 和 r 不是找到的公共祖先，而是递归调用的返回值，它们可能表示：
	// 1:找到了目标节点之一（p 或 q）
	// 2:找到了最近公共祖先
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	// 关键判断
	if l != nil && r != nil {
		// 这意味着 p 和 q 分别在当前节点的左右子树中
		return root // 当前节点就是LCA
	}

	// 只在一侧找到，返回非空的那一侧
	if l != nil {
		return l
	}
	return r
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 基础情况
	if root == nil || root == p || root == q {
		return root
	}
	if q.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
