package main

import (
	"sort"
)

func main() {

}

func maximumGroups(grades []int) int {
	sort.Ints(grades)
	pre := 1
	ans := 0
	n := len(grades)
	i := 0
	for i < n {
		if i+pre < n {
			ans++
			pre++
			i = i + pre
		} else {
			break
		}
	}
	if i < n {
		ans++
	}
	return ans
}
