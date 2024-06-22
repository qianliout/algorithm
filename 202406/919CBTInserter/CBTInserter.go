package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	cb := Constructor(root)
	cb.Insert(2)
	cb.Insert(3)
	cb.Insert(4)
	cb.Insert(5)
	fmt.Println(cb.ROOT.Val)
}

type CBTInserter struct {
	ROOT      *TreeNode
	Candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	q := []*TreeNode{root}
	candidate := make([]*TreeNode, 0)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}

	return CBTInserter{
		ROOT:      root,
		Candidate: candidate,
	}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{Val: val}
	can := this.Candidate[0]
	if can.Left == nil {
		can.Left = node
	} else if can.Right == nil {
		can.Right = node
		this.Candidate = this.Candidate[1:]
	}
	this.Candidate = append(this.Candidate, node)
	return can.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.ROOT
}
