package main

import (
	"sort"
	"strings"
)

func main() {

}

func findMaxForm(strs []string, m int, n int) int {

	nums := make([]Node, len(strs))
	for i, ch := range strs {
		nums[i] = Node{strings.Count(ch, "1"), strings.Count(ch, "0")}
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i].one == nums[j].one {
			return nums[i].zero < nums[j].zero
		}
		return nums[i].one < nums[j].one
	})
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if
	}

}

type Node struct {
	one, zero int
}
