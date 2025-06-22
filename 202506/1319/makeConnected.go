package main

import "fmt"

func main() {
	// 测试用例1：基本例子
	n1 := 4
	connections1 := [][]int{{0, 1}, {0, 2}, {1, 2}}
	fmt.Printf("测试1: n=%d, connections=%v\n", n1, connections1)
	result1 := makeConnected(n1, connections1)
	fmt.Printf("结果: %d\n\n", result1)

	// 测试用例2：不可能的情况
	n2 := 6
	connections2 := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}
	fmt.Printf("测试2: n=%d, connections=%v\n", n2, connections2)
	result2 := makeConnected(n2, connections2)
	fmt.Printf("结果: %d\n\n", result2)

	// 详细分析
	fmt.Println("=== 详细分析过程 ===")
	makeConnectedDetailed(n1, connections1)
}

// 简单并查集实现
type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, count: n}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) bool {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return false // 已经连通
	}
	uf.parent[px] = py
	uf.count--
	return true
}

func (uf *UnionFind) isConnected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func makeConnected(n int, connections [][]int) int {
	/*
		🎯 核心思想：图论 + 并查集

		关键洞察：
		1. 连通分量：当前网络被分成几个独立的部分
		2. 冗余线缆：同一连通分量内的多余线缆
		3. 最少操作：连通分量数 - 1

		为什么这样计算？
		- 要连通k个分量，需要k-1条线缆
		- 冗余线缆可以重新利用
		- 如果冗余线缆不够，则无法连通
	*/

	uf := NewUnionFind(n)
	redundant := 0 // 冗余线缆数量

	// 统计冗余线缆
	for _, connection := range connections {
		x, y := connection[0], connection[1]
		if uf.isConnected(x, y) {
			// 如果两个节点已经连通，这条线缆就是冗余的
			redundant++
		}
		uf.union(x, y)
	}

	components := uf.count   // 连通分量数量
	needed := components - 1 // 需要的线缆数量

	// 判断是否可能连通
	if needed > redundant {
		return -1 // 冗余线缆不够
	}

	return needed // 返回最少操作次数
}

// 详细分析函数 - 不依赖外部并查集，自己实现
func makeConnectedDetailed(n int, connections [][]int) {
	fmt.Printf("分析网络: %d台计算机, %d条线缆\n", n, len(connections))
	fmt.Printf("线缆连接: %v\n\n", connections)

	// 简单并查集实现
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(x, y int) bool {
		px, py := find(x), find(y)
		if px == py {
			return false // 已经连通
		}
		parent[px] = py
		return true
	}

	redundant := 0
	fmt.Println("处理线缆过程:")

	for i, conn := range connections {
		x, y := conn[0], conn[1]
		fmt.Printf("线缆%d: 连接 %d-%d", i+1, x, y)

		if !union(x, y) {
			redundant++
			fmt.Printf(" -> 冗余线缆 (已连通)")
		} else {
			fmt.Printf(" -> 新连接")
		}
		fmt.Println()
	}

	// 统计连通分量
	components := 0
	componentMap := make(map[int][]int)

	for i := 0; i < n; i++ {
		root := find(i)
		if _, exists := componentMap[root]; !exists {
			components++
		}
		componentMap[root] = append(componentMap[root], i)
	}

	fmt.Printf("\n连通分量分析:\n")
	fmt.Printf("总连通分量数: %d\n", components)

	compNum := 1
	for root, nodes := range componentMap {
		fmt.Printf("分量%d (根节点%d): %v\n", compNum, root, nodes)
		compNum++
	}

	needed := components - 1
	fmt.Printf("\n计算过程:\n")
	fmt.Printf("冗余线缆数: %d\n", redundant)
	fmt.Printf("连通分量数: %d\n", components)
	fmt.Printf("需要线缆数: %d - 1 = %d\n", components, needed)

	if needed > redundant {
		fmt.Printf("结果: -1 (冗余线缆不够，需要%d条但只有%d条)\n", needed, redundant)
	} else {
		fmt.Printf("结果: %d (最少操作次数)\n", needed)
	}

	// 可视化解释
	fmt.Printf("\n💡 为什么是这个结果？\n")
	fmt.Printf("1. 要连通%d个分量，需要%d条线缆\n", components, needed)
	fmt.Printf("2. 我们有%d条冗余线缆可以重新利用\n", redundant)
	if needed <= redundant {
		fmt.Printf("3. 冗余线缆足够，所以需要%d次操作\n", needed)
	} else {
		fmt.Printf("3. 冗余线缆不够，无法完成连通\n")
	}
}

/*
用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。
网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。
给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。
*/
