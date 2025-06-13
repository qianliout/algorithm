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
	ans := ser(root)
	fmt.Println(ans)
	return ans
}

func ser(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	left := ser(root.Left)
	right := ser(root.Right)
	ans := []string{fmt.Sprintf("%d", root.Val), left, right}
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	qu := strings.Split(data, ",")
	root := des(&qu)
	return root
}

// 这里一定要传指针，不然就不能得到正确的结果
func des(qu *[]string) *TreeNode {
	if len(*qu) == 0 {
		return nil
	}
	fir := (*qu)[0]
	*qu = (*qu)[1:]
	if fir == "null" {
		return nil
	}
	i, _ := strconv.Atoi(fir)
	root := &TreeNode{Val: i}
	left := des(qu)
	right := des(qu)
	root.Left = left
	root.Right = right
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
