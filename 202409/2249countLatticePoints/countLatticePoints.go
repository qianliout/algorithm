package main

import (
	"sort"
)

func main() {

}

func countLatticePoints(circles [][]int) int {
	// 可以先按半径排序，但是这里是否排序都不影响结果
	sort.Slice(circles, func(i, j int) bool {
		return circles[i][2] > circles[j][2]
	})
	ans := 0
	for i := 0; i <= 200; i++ {
		for j := 0; j <= 200; j++ {
			for _, ch := range circles {
				x, y, r := ch[0], ch[1], ch[2]
				// 判断是否在其中
				if (i-x)*(i-x)+(j-y)*(j-y) <= r*r {
					ans++
					break
				}
			}
		}
	}
	return ans
}
