package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	root := recoverFromPreorder("1-2--3---4-5--6---7")
	fmt.Println(root.Val)
}

func recoverFromPreorder(traversal string) *TreeNode {
	return nil
}
