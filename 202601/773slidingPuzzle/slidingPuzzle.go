package main

import (
	"fmt"
	"strings"
)

// 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示。
// 一次 移动 定义为选择 0 与一个相邻的数字（上下左右）进行交换.
// 最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
// 给出一个谜板的初始状态 board ，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
func slidingPuzzle(board [][]int) int {
	target := "123450"
	var sb strings.Builder
	for _, row := range board {
		for _, v := range row {
			sb.WriteString(fmt.Sprintf("%d", v))
		}
	}
	start := sb.String()

	if start == target {
		return 0
	}

	// 0 1 2
	// 3 4 5
	// neighbors 映射每个位置可以交换的相邻位置索引
	neighbors := [][]int{
		{1, 3},    // 0
		{0, 2, 4}, // 1
		{1, 5},    // 2
		{0, 4},    // 3
		{1, 3, 5}, // 4
		{2, 4},    // 5
	}

	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true

	step := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == target {
				return step
			}

			// 找到 '0' 的位置
			zeroIdx := strings.Index(curr, "0")

			// 遍历相邻位置
			for _, neighborIdx := range neighbors[zeroIdx] {
				newBoardBytes := []byte(curr)
				// 交换
				newBoardBytes[zeroIdx], newBoardBytes[neighborIdx] = newBoardBytes[neighborIdx], newBoardBytes[zeroIdx]
				newBoardStr := string(newBoardBytes)

				if !visited[newBoardStr] {
					if newBoardStr == target {
						return step + 1
					}
					visited[newBoardStr] = true
					queue = append(queue, newBoardStr)
				}
			}
		}
		step++
	}

	return -1
}
