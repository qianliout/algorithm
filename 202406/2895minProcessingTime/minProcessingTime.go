package main

import (
	"sort"
)

func main() {

}

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(processorTime)
	sort.Slice(tasks, func(i, j int) bool { return tasks[i] > tasks[j] })
	ans := 0
	for i, j := 0, 0; i < len(tasks); i, j = i+4, j+1 {
		ans = max(ans, processorTime[j]+Max(tasks[i:i+4]))
	}

	return ans
}

func Max(task []int) int {
	ans := 0
	for _, ch := range task {
		ans = max(ans, ch)
	}
	return ans
}
