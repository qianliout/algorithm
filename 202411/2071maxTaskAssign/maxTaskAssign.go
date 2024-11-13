package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTaskAssign([]int{3, 2, 1}, []int{0, 3, 3}, 1, 1))
	fmt.Println(maxTaskAssign([]int{5, 9, 8, 5, 9}, []int{1, 6, 4, 2, 6}, 1, 5))
}

func maxTaskAssign1(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	m, n := len(tasks), len(workers)
	var check func(i int) bool

	check = func(mid int) bool {
		if mid > min(m, n) {
			return false
		}
		if mid < 0 {
			return false
		}
		subTask := tasks[:mid]
		subWork := workers[n-mid:]
		used := 0
		validTask := list.New()
		taskIdx := 0
		for j := 0; j < mid; j++ {
			// 将所有可以分配的任务加入到有效任务队列中
			for taskIdx < len(subTask) && subTask[taskIdx] <= subWork[j]+strength {
				validTask.PushBack(subTask[taskIdx])
				taskIdx++
			}
			// 如果没有有效任务，标记为失败
			if validTask.Len() == 0 {
				return false
			}
			// 如果任务可以直接分配，从队列前端移除任务
			if validTask.Front().Value.(int) <= subWork[j] {
				validTask.Remove(validTask.Front())
				continue
			}
			// 如果需要药丸且药丸数量足够，从队列后端移除任务
			if used >= pills {
				return false
			}
			used++
			validTask.Remove(validTask.Back())
		}

		return true
	}

	cnt := min(m, n) + 1
	left, right := 0, cnt
	for left < right {
		// 找最多，也就是右端点
		mid := left + (right-left+1)/2
		if mid < cnt && check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)
	m, n := len(tasks), len(workers)
	var check func(i int) bool

	check = func(mid int) bool {
		if mid > min(m, n) {
			return false
		}
		if mid < 0 {
			return false
		}
		used := 0
		validTask := list.New()
		taskIdx := 0
		for j := n - mid; j < n; j++ {
			// 将所有可以分配的任务加入到有效任务队列中
			for taskIdx < m && tasks[taskIdx] <= workers[j]+strength {
				validTask.PushBack(tasks[taskIdx])
				taskIdx++
			}
			// 如果没有有效任务，标记为失败
			if validTask.Len() == 0 {
				return false
			}
			// 如果任务可以直接分配，从队列前端移除任务
			if validTask.Front().Value.(int) <= workers[j] {
				validTask.Remove(validTask.Front())
				continue
			}
			// 如果需要药丸且药丸数量足够，从队列后端移除任务
			if used >= pills {
				return false
			}
			used++
			validTask.Remove(validTask.Back())
		}
		return true
	}

	cnt := min(m, n) + 1
	left, right := 0, cnt
	for left < right {
		// 找最多，也就是右端点
		mid := left + (right-left+1)/2
		if mid < cnt && check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
