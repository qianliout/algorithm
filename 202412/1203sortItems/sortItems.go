package main

func main() {

}

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
