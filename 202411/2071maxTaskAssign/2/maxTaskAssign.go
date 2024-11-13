package main

import (
	"container/list"
	"sort"
)

type Solution struct{}

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	sort.Ints(tasks)
	sort.Ints(workers)

	start := 0
	end := min(len(tasks), len(workers)) + 1

	for start+1 < end {
		mid := (start + end) / 2

		if canAssign(tasks, workers, pills, strength, mid) {
			start = mid
		} else {
			end = mid
		}
	}

	return start
}

func canAssign(tasks []int, workers []int, pills int, strength int, m int) bool {
	i2 := 0
	p := pills
	fail := false
	validTasks := list.New()

	for j := len(workers) - m; j < len(workers); j++ {
		w := workers[j]

		// 这一步是关键
		for i2 < m && tasks[i2] <= w+strength {
			validTasks.PushBack(tasks[i2])
			i2++
		}

		if validTasks.Len() == 0 {
			fail = true
			break
		}

		if validTasks.Front().Value.(int) <= w {
			// No need for pill
			validTasks.Remove(validTasks.Front())
		} else {
			if p == 0 {
				fail = true
				break
			}
			p--
			validTasks.Remove(validTasks.Back())
		}
	}

	return !fail
}
