package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	root := &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
	fmt.Println(smallestFromLeaf(root))
}

// 这一题目很容易出错的地方是：
// 如果root.Left==nil，root.Right!=nil,那么只取 root 时不认为已经走到了叶子节点
func smallestFromLeaf(root *TreeNode) string {
	ans := ""
	if root == nil {
		return ans
	}
	var dfs func(root *TreeNode, path []byte)
	dfs = func(root *TreeNode, path []byte) {

		path = append(path, byte('a'+root.Val))

		if root.Left == nil && root.Right == nil {
			tem := append([]byte{}, path...)
			reverse(tem)
			if ans == "" || string(tem) < ans {
				ans = string(tem)
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
	}
	dfs(root, []byte{})
	return ans
}

func reverse(data []byte) {
	le, ri := 0, len(data)-1
	for le < ri {
		data[le], data[ri] = data[ri], data[le]
		le++
		ri--
	}
}
