package main

import (
	"fmt"
)

func main() {
	gr := []int{-1, -1, 1, 0, 0, 1, 0, -1}
	be := [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}
	fmt.Println(sortItems(8, 2, gr, be))
}

/*
有 n 个项目，每个项目或者不属于任何小组，或者属于 m 个小组之一。group[i] 表示第 i 个项目所属的小组，如果第 i 个项目不属于任何小组，
则 group[i] 等于 -1。项目和小组都是从零开始编号的。可能存在小组不负责任何项目，即没有任何项目属于这个小组。
请你帮忙按要求安排这些项目的进度，并返回排序后的项目列表：
	同一小组的项目，排序后在列表中彼此相邻。
	项目之间存在一定的依赖关系，我们用一个列表 beforeItems 来表示，其中 beforeItems[i] 表示在进行第 i 个项目前（位于第 i 个项目左侧）应该完成的所有项目。
如果存在多个解决方案，只需要返回其中任意一个即可。如果没有合适的解决方案，就请返回一个 空列表 。
*/

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	for i, ch := range group {
		if ch == -1 {
			group[i] = m
			m++
		}
	}
	// 建图，本次使用邻接表
	// 这里的 m 值是一个易错点，上面对元数据做了处理，没有对应组的，我们认为他自成一个组，且是比原组数大的数，所以到这里后，组的个数应该是变动之后的数
	groupObj := make([][]int, m) // 要理解邻接表的保存数据的方法
	itemObj := make([][]int, n)
	// 入度，理解成项目的依赖
	groupIn := make([]int, m)
	itemIn := make([]int, n)
	// 对项目建图
	for ite := 0; ite < n; ite++ {
		for _, breIte := range beforeItems[ite] {
			itemObj[breIte] = append(itemObj[breIte], ite)
			itemIn[ite]++
		}
	}
	// 对组建图
	for i, curGroup := range group {
		befItem := beforeItems[i]
		for _, bi := range befItem {
			befGroup := group[bi]
			// 自已不能依赖自己
			if befGroup != curGroup { // 这里的判断有点难想，和上面对数据的处理有关
				groupObj[befGroup] = append(groupObj[befGroup], curGroup)
				groupIn[curGroup]++
			}
		}
	}
	groupSort := topologicalSort(groupObj, groupIn, m)
	itemSort := topologicalSort(itemObj, itemIn, n)
	if len(groupSort) == 0 || len(itemSort) == 0 {
		return []int{}
	}

	group2item := make([][]int, m)

	for _, it := range itemSort {
		g := group[it]
		group2item[g] = append(group2item[g], it)
	}
	res := make([]int, 0)
	for _, gi := range groupSort {
		ints := group2item[gi]
		res = append(res, ints...)
	}
	return res
}

func topologicalSort(adj [][]int, in []int, m int) []int {
	res := make([]int, 0)
	queue := make([]int, 0)
	for k, v := range in {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	for len(queue) > 0 {
		fir := queue[0]
		queue = queue[1:]
		res = append(res, fir)
		for _, ad := range adj[fir] {
			in[ad]--
			if in[ad] == 0 {
				queue = append(queue, ad)
			}
		}
	}
	if len(res) == m {
		return res
	}
	return []int{}
}
