package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	root := recoverFromPreorder("1-2--3---4-5--6---7")
	fmt.Println(root.Val)
}

// 这种方式有错，原因是，同层节点时，挂载的父节点会不一样
func recoverFromPreorder(traversal string) *TreeNode {
	return nil
}
