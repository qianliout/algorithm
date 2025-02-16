package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotateGrid([][]int{{40, 10}, {30, 20}}, 1)) // [10,20],[40,30]]
	fmt.Println(rotateGrid([][]int{{10, 1, 4, 8}, {6, 6, 3, 10}, {7, 4, 7, 10}, {1, 10, 6, 1}, {2, 1, 1, 10}, {3, 8, 9, 2}, {7, 1, 10, 10}, {7, 1, 4, 9}, {2, 2, 4, 2}, {10, 7, 5, 10}}, 1))
}

func rotateGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	le, ri, top, down, cnt := 0, m-1, 0, n-1, 0

	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}
	for cnt <= m*n {
		level := make([]index, 0)
		nums := make([]int, 0)

		// 向下
		for t := top; t <= down; t++ {
			level = append(level, index{i: t, j: le})
			nums = append(nums, grid[t][le])
		}
		le++
		// 向右
		for l := le; l <= ri; l++ {
			level = append(level, index{i: down, j: l})
			nums = append(nums, grid[down][l])
		}
		down--
		// 向上
		for d := down; d >= top; d-- {
			level = append(level, index{i: d, j: ri})
			nums = append(nums, grid[d][ri])
		}
		ri--
		// 向左
		for r := ri; r >= le; r-- {
			level = append(level, index{i: top, j: r})
			nums = append(nums, grid[top][r])
		}
		top++
		// 旋转
		cn := len(nums)
		if cn == 0 {
			break
		}

		b := k % cn
		// 这里是这一题的关键
		nums2 := append(nums[cn-b:], nums[:cn-b]...)

		for i := 0; i < cn; i++ {
			ans[level[i].i][level[i].j] = nums2[i]
		}

		cnt += cn
		if cnt >= m*n {
			break
		}
	}
	return ans
}

type index struct {
	i, j int
}
