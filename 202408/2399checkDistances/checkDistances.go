package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func checkDistances(s string, distance []int) bool {
	dis := make([]int, 26)
	for i := 0; i < 26; i++ {
		dis[i] = -1
	}
	for i, ch := range s {
		idx := int(ch) - int('a')
		if dis[idx] == -1 {
			dis[idx] = i
		} else {
			dis[idx] = i - dis[idx] - 1
		}
	}
	for i := 0; i < 26; i++ {
		if dis[i] == -1 {
			continue
		}
		if dis[i] != distance[i] {
			return false
		}
	}
	return true
}
