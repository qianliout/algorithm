package main

import (
	"fmt"
)

func main() {

}

func isPathCrossing(path string) bool {
	// N'、'S'、'E' 或者 'W'，分别表示向北、向南、向东、向西移动一个单位。
	dirs := map[string][]int{
		"N": []int{1, 0},
		"S": []int{-1, 0},
		"E": []int{0, 1},
		"W": []int{0, -1},
	}
	point := map[string]bool{}
	x, y := 0, 0
	point[fmt.Sprintf("%d-%d", x, y)] = true
	for _, ch := range path {
		dir := dirs[string(ch)]
		x, y = x+dir[0], y+dir[1]
		key := fmt.Sprintf("%d-%d", x, y)
		if point[key] {
			return true
		}
		point[key] = true
	}
	return false
}
