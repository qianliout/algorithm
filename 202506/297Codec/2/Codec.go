package main

import (
	"fmt"
	"strconv"
	"strings"

	. "outback/algorithm/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	code := Constructor()
	ans := code.serialize(root)
	node := code.deserialize(ans)
	fmt.Println(node.Val)
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
	ans := make([]string, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no == nil {
				ans = append(ans, "null")
			} else {
				ans = append(ans, fmt.Sprintf("%d", no.Val))
				lev = append(lev, no.Left)
				lev = append(lev, no.Right)
			}
		}
		queue = lev
	}
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	root := &TreeNode{}
	que := []*TreeNode{root}
	for len(ss) > 0 {
		fir := ss[0]
		ss = ss[1:]
		qu := que[0]
		que = que[1:]
		atoi, err := strconv.Atoi(fir)
		if err == nil {
			if qu == nil {
				qu = &TreeNode{}
			}

			qu.Val = atoi
			que = append(que, qu.Left)
			que = append(que, qu.Right)
		}
	}
	return root
}
