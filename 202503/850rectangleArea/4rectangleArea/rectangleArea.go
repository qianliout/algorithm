package main

import (
	"sort"
)

func main() {

}

func rectangleArea(rectangles [][]int) int {
	MOD := int64(1e9 + 7)
	xs := []int{}
	for _, info := range rectangles {
		xs = append(xs, info[0])
		xs = append(xs, info[2])
	}
	sort.Ints(xs)

	ans := int64(0)
	for i := 1; i < len(xs); i++ {
		a, b, length := xs[i-1], xs[i], xs[i]-xs[i-1]
		if length == 0 {
			continue
		}
		ys := [][]int{}
		for _, info := range rectangles {
			if info[0] <= a && b <= info[2] {
				ys = append(ys, []int{info[1], info[3]})
			}
		}
		sort.Slice(ys, func(i, j int) bool {
			if ys[i][0] != ys[j][0] {
				return ys[i][0]-ys[j][0] < 0
			}
			return ys[i][1]-ys[j][1] < 0
		})
		// 求一组线段的并集
		total, l, r := int64(0), -1, -1
		for _, cur := range ys {
			if cur[0] > r {
				total += int64(r - l)
				l, r = cur[0], cur[1]
			} else if cur[1] > r {
				r = cur[1]
			}
		}
		total += int64(r - l)
		ans += total * int64(length)
		ans %= MOD
	}
	return int(ans)
}

// https://leetcode.cn/problems/rectangle-area-ii/solutions/1826992/gong-shui-san-xie-by-ac_oier-9r36/
// 宫业三水的题解
