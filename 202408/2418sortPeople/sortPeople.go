package main

import (
	"sort"
)

func main() {

}

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	p := make([]pair, n)
	for i := 0; i < n; i++ {
		p[i] = pair{names[i], heights[i]}
	}
	sort.Slice(p, func(i, j int) bool { return p[i].height > p[j].height })
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = p[i].name
	}
	return ans
}

type pair struct {
	name   string
	height int
}
