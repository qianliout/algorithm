package main

import (
	"sort"
)

func main() {

}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		if tiles[i][0] < tiles[j][0] {
			return true
		} else if tiles[i][0] > tiles[j][0] {
			return false
		}
		return tiles[i][1] < tiles[j][1]
	})
	le, ri := 0, 0
	n := len(tiles)
	ans := 0
	cov := 0
	for le <= ri && ri < n {
		rb := tiles[le][0] + carpetLen - 1 // 最长边边界
		// 能完全覆盖 ri 这个区间
		if rb >= tiles[ri][1] {
			cov += tiles[ri][1] - tiles[ri][0] + 1
			ri++
			ans = max(ans, cov)
		} else {
			// 只能覆盖部分区间，这里就要先算结果，再把 le向右移
			if rb >= tiles[ri][0] {
				ans = max(ans, cov+rb-tiles[ri][0]+1)
			}
			// 调整到下一个区间开头
			cov -= tiles[le][1] - tiles[le][0] + 1
			le++
		}
	}
	return ans
}
