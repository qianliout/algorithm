package main

import (
	"sort"
)

func main() {

}

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1]-tasks[i][0] <= tasks[j][1]-tasks[j][0] })
	res := 0
	for _, ch := range tasks {
		res = max(ch[1], res+ch[0])
	}
	return res
}
