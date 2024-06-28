package main

import (
	"fmt"
)

func main() {
	fmt.Println(isRobotBounded("GL"))
}

// 没有做对
func isRobotBounded(instructions string) bool {

	dirs := []int{0, 1, 0, -1, 0}
	k, x, y := 0, 0, 1
	pos := []int{0, 0}
	obs := make(map[string]bool)
	obs[fmt.Sprintf("%d-%d", 0, 0)] = true
	for _, c := range instructions {
		if c == 'L' {
			k = (k + 3) % 4
		} else if c == 'R' {
			k = (k + 1) % 4
		} else if c == 'G' {
			x = dirs[k]
			y = dirs[k+1]
			pos = []int{pos[0] + x, pos[1] + y}

			key := fmt.Sprintf("%d-%d", pos[0], pos[1])
			if obs[key] {
				return true
			}
			obs[key] = true
		}
	}
	return false
}
