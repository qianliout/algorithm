package main

import (
	"container/heap"
	"fmt"
	"sort"

	. "outback/algorithm/common/commonHeap"
)

func main() {
	fmt.Println(kthSmallest([][]int{{1, 3, 11}, {2, 4, 6}}, 5))
}

func kthSmallest2(mat [][]int, k int) int {
	hp := make(MaxHeap, 0)
	for _, ch := range mat {
		for _, nu := range ch {
			if hp.Len() < k {
				heap.Push(&hp, nu)
				continue
			}
			if hp[0] >= nu {
				continue
			}
			heap.Pop(&hp)
			heap.Push(&hp, nu)

		}
	}
	return hp[0]
}

func kthSmallest3(mat [][]int, k int) int {
	a := []int{0}
	for _, row := range mat {
		b := make([]int, 0)
		for _, x := range row {
			// 这里有循环，如果上面不给a加一个0，这里就要特殊处理
			for _, y := range a {
				b = append(b, x+y)
			}
		}
		sort.Ints(b)
		if len(b) > k {
			b = b[:k]
		}
		a = b
	}
	return a[len(a)-1]
}

func kthSmallest(mat [][]int, k int) int {
	a := []int{}
	for _, row := range mat {
		b := make([]int, 0)
		for _, x := range row {
			if len(a) == 0 {
				// 这里有循环，如果上面不给a加一个0，这里就要特殊处理
				b = append(b, x)
			} else {
				for _, y := range a {
					b = append(b, x+y)
				}
			}
		}
		sort.Ints(b)
		if len(b) > k {
			b = b[:k]
		}
		a = b
	}
	return a[len(a)-1]
}
