package main

import (
	"fmt"
)

func main() {
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))
}

func robotSim(commands []int, obstacles [][]int) int {
	dirs := []int{0, 1, 0, -1, 0}
	k, x, y, ans := 0, 0, 0, 0
	obs := make(map[string]bool)
	for _, ch := range obstacles {
		key := fmt.Sprintf("%d-%d", ch[0], ch[1])
		obs[key] = true
	}
	for _, c := range commands {
		if c == -2 {
			k = (k + 3) % 4
		} else if c == -1 {
			k = (k + 1) % 4
		} else {
			for i := 0; i < c; i++ {
				nx, ny := x+dirs[k], y+dirs[k+1]
				key := fmt.Sprintf("%d-%d", nx, ny)
				if obs[key] {
					break
				}
				// 这里一定要先检测障碍物，之后再更新 x,y
				x, y = nx, ny
				ans = max(ans, x*x+y*y)
			}
		}
	}
	return ans
}
