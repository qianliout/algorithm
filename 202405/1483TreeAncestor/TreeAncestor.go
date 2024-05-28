package main

import (
	"fmt"
	"math/bits"
)

func main() {

	fmt.Println(bits.Len(7))
	fmt.Println(bits.Len(8))
}

type TreeAncestor struct {
	Data   [][]int
	N      int
	Parent []int
}

func Constructor(n int, parent []int) TreeAncestor {
	dep := bits.Len(uint(n)) // n有多少位,也就是说要处理多少层,极端情况下 parent 是一个链表，会跳多少次

	// pa[i][j] 表示 i 这个元素的，第 j 层，也就是第 j 个祖先节点
	pa := make([][]int, n)
	for i, p := range parent {
		pa[i] = make([]int, dep)
		// 第0个 祖先节点就是父节点
		pa[i][0] = p
	}

	for i := 0; i < dep-1; i++ {
		for x := 0; x < n; x++ {
			p := pa[x][i]
			if p != -1 {
				pa[x][i+1] = pa[p][i]
			} else {
				pa[x][i+1] = -1
			}
		}
	}
	return TreeAncestor{Data: pa, N: n, Parent: parent}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	dep := bits.Len(uint(this.N))
	for i := 0; i < dep; i++ {
		if k>>i&1 > 0 {
			node = this.Data[node][i]
			if node < 0 {
				break
			}
		}
	}
	return node
}
