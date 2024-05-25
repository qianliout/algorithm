package main

import (
	"fmt"
)

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {}, {0, 3}, {1}}))
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	vis := make([]bool, n)
	dfs(rooms, 0, vis)
	if check(vis) {
		return true
	}
	return false
}

func check(vis []bool) bool {
	for k, v := range vis {
		if !v && k != 0 {
			return false
		}
	}
	return true
}

func dfs(rooms [][]int, start int, vis []bool) {
	sec := rooms[start]
	for _, ch := range sec {
		if vis[ch] {
			continue
		}
		vis[ch] = true
		dfs(rooms, ch, vis)
	}
}
