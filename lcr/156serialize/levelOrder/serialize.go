package main

import (
	"fmt"
	. "outback/algorithm/common/treenode"
	"strconv"
	"strings"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	c := Constructor()
	ss := c.serialize(root)
	fmt.Println(ss)
	node := c.deserialize(ss)
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

	path := make([]string, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, node := range queue {
			if node == nil {
				path = append(path, "nil")
				continue
			}

			path = append(path, fmt.Sprintf("%d", node.Val))
			lev = append(lev, node.Left)
			lev = append(lev, node.Right)
		}
		queue = lev
	}

	return strings.Join(path, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	if len(ss) == 0 || ss[0] == "nil" {
		return nil
	}
	queue := make([]*TreeNode, 0)
	root := this.node(ss[0])
	queue = append(queue, root)
	idx := 1
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, node := range queue {
			// 这里不怕 idx 越界的原因是上面序列化时已经是满树
			node.Left = this.node(ss[idx])
			node.Right = this.node(ss[idx+1])
			idx += 2
			if node.Left != nil {
				lev = append(lev, node.Left)
			}
			if node.Right != nil {
				lev = append(lev, node.Right)
			}
		}
		queue = lev
	}

	return root
}

func (this *Codec) node(s string) *TreeNode {
	if s == "nil" {
		return nil
	}
	n, _ := strconv.Atoi(s)
	return &TreeNode{Val: n}
}
