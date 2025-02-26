package main

import (
	"fmt"
	. "outback/algorithm/common/treenode"
	"strconv"
	"strings"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	path := make([]string, 0)
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			path = append(path, "nil")
			return
		}
		path = append(path, fmt.Sprintf("%d", root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(path, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	return this.des(&ss)
}

func (this *Codec) des(data *[]string) *TreeNode {
	if len(*data) == 0 {
		return nil
	}
	if (*data)[0] == "nil" {
		*data = (*data)[1:]
		return nil
	}
	n, _ := strconv.Atoi((*data)[0])
	root := &TreeNode{Val: n}
	*data = (*data)[1:]
	root.Left = this.des(data)
	root.Right = this.des(data)
	return root
}
