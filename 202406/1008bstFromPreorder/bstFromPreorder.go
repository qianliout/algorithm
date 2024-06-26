package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}
	ri := find(preorder)
	left := bstFromPreorder(preorder[1:ri])
	right := bstFromPreorder(preorder[ri:])
	root.Left = left
	root.Right = right
	return root
}

// 找右子树的启点
// 对于二叉搜索树的前序遍历，root 在第0号位置，nums 中第一个大于 nums[0]的元素就是 right 的启点
func find(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			return i
		}
	}
	return len(nums)
}
