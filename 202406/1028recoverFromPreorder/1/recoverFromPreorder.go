package main

import (
	"fmt"

	. "outback/algorithm/common/treenode"
)

func main() {
	parse("1-2--3---4-5--6---7")
	root := recoverFromPreorder("1-2--3---4-5--6---7")
	fmt.Println(root.Val)
}

// 这种方式有错，原因是，同层节点时，挂载的父节点会不一样
func recoverFromPreorder(traversal string) *TreeNode {
	res := parse(traversal)
	if len(res[0]) == 0 {
		return nil
	}
	dep := 0
	root := &TreeNode{Val: res[0][0].Val}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		dep++
		lev := make([]*TreeNode, 0)
		next := res[dep]
		for _, no := range queue {
			if len(next) == 0 {
				break
			}
			if len(next) > 0 {
				no.Left = &TreeNode{Val: next[0].Val}
				lev = append(lev, no.Left)
				next = next[1:]
			}
			if len(next) > 0 {
				no.Right = &TreeNode{Val: next[0].Val}
				lev = append(lev, no.Right)
				next = next[1:]
			}
		}
		queue = lev
	}

	return root
}

func parse(tra string) map[int][]pair {
	res := make(map[int][]pair)
	dep := 0
	num := 0
	for i := 0; i < len(tra); i++ {
		if tra[i] == '-' {
			if i > 0 && tra[i-1] != '-' {
				res[dep] = append(res[dep], pair{Dep: dep, Val: num})
				num = 0
				dep = 0
			}
			dep++
			continue
		} else {
			num = num*10 + int(tra[i]) - int('0')
		}
	}
	if num > 0 {
		res[dep] = append(res[dep], pair{Dep: dep, Val: num})
	}

	return res
}

type pair struct {
	Val int
	Dep int
}
