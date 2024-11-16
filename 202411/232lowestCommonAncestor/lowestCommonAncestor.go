package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}
	le := lowestCommonAncestor(root.Left, p, q)
	ri := lowestCommonAncestor(root.Right, p, q)
	// 这里是最不好理解的，在左边也找到了，在右边也找到了，那么最近公共祖先就是 root,这个怎么理解呢，
	// 按正常的理解，这里不可能在两边都找到一个公共祖先，为啥能两边都可以不是 nil 呢，原因就是第一行，
	// if root==p || root==q  返回是 root

	if le != nil && ri != nil {
		return root
	}
	if le == nil && ri == nil {
		return nil
	}
	if le == nil {
		return ri
	}
	return le
}
