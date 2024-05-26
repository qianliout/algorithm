package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}))
}

/*
首先 idx 是用来对边进行编号的，然后对存图用到的几个数组作简单解释：

	he 数组：存储是某个节点所对应的边的集合（链表）的头结点；
	e  数组：用于访问某一条边指向的节点；
	ne 数组：用于是以链表的形式进行存边，该数组就是用于找到下一条边；
	w  数组：用于记录某条边的权重为多少。

int[] he = new int[N], e = new int[M], ne = new int[M], w = new int[M];
int idx;

	void add(int a, int b, int c) {
	    e[idx] = b;
	    ne[idx] = he[a];
	    he[a] = idx;
	    w[idx] = c;
	    idx++;
	}
*/
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	in := make([][]int, n)
	out := make([]int, n)
	for i, ch := range graph {
		out[i] = len(ch)
		for _, j := range ch {
			if j != i {
				in[j] = append(in[j], i)
			}
		}
	}
	// 找到出度为0的点
	q := make([]int, 0)
	for i, c := range out {
		if c == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, 0)
	for len(q) > 0 {
		fir := q[0]
		ans = append(ans, fir)
		q = q[1:]
		// 找到以这些点作为出度的点
		for _, j := range in[fir] {
			out[j]--
			if out[j] == 0 {
				q = append(q, j)
			}
		}
	}

	sort.Ints(ans)

	return ans
}
