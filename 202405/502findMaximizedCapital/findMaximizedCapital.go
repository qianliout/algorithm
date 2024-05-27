package main

import (
	"container/heap"
	"fmt"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}

func findMaximizedCapital2(k int, w int, profits []int, capital []int) int {
	type pair struct{ p, c int }
	pc := make([]pair, len(profits))
	for i := range profits {
		pc[i] = pair{profits[i], capital[i]}
	}
	sort.Slice(pc, func(i, j int) bool { return pc[i].c < pc[j].c })

	queue := make([]pair, 0)

	used := make([]bool, len(pc))
	for k > 0 {
		// 每次决策前，将所有的启动资金不超过 www 的任务加入优先队列（根据利润排序的大根堆），然后从优先队列（根据利润排序的大根堆），将利润累加到 www；
		for i := 0; i < len(pc); i++ {
			if pc[i].c <= w && !used[i] {
				used[i] = true
				queue = append(queue, pc[i])
			}
		}
		// 每次都排序会超时，
		sort.Slice(queue, func(i, j int) bool { return queue[i].p > queue[j].p })

		if len(queue) == 0 {
			break
		}
		w += queue[0].p
		queue = queue[1:]
		k--
	}
	return w
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// 必须用两个优先队列，不然会timeout

	pc := make(PriorityQueue, 0)

	for i := range profits {
		// 默认是从大小到排序，所以这里转一下负数
		heap.Push(&pc, &IntItem{Value: profits[i], Priority: -capital[i]})
	}

	queue := make(PriorityQueue, 0)

	for k > 0 {
		// 每次决策前，将所有的启动资金不超过 www 的任务加入优先队列（根据利润排序的大根堆），然后从优先队列（根据利润排序的大根堆），将利润累加到 www；
		for len(pc) > 0 {
			pop := heap.Pop(&pc).(*IntItem)
			if -pop.Priority > w {
				heap.Push(&pc, pop)
				break
			}
			heap.Push(&queue, &IntItem{Value: -pop.Priority, Priority: pop.Value})
		}
		if len(queue) == 0 {
			break
		}
		pop := heap.Pop(&queue).(*IntItem)
		w += pop.Priority
		k--
	}
	return w
}
