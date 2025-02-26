package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 这里是最不好理解的，在左边也找到了，在右边也找到了，那么最近公共祖先就是 root,这个怎么理解呢，
	// 按正常的理解，这里不可能在两边都找到一个公共祖先，为啥能两边都可以不是 nil 呢，原因就是第一行，
	// if root==p || root==q  返回是 root

	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	return left
}

// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	a := n / 2
	b := myPow(x, a)
	ans := b * b
	if n%2 == 1 {
		ans = ans * x
	}
	return ans
}
