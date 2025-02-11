package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxDistance("NWSE", 1))
	fmt.Println(maxDistance("NSWWEW", 3))
}

func maxDistance(s string, k int) int {
	ans, x, y := 0, 0, 0
	for i, c := range s {
		switch byte(c) {
		case 'N':
			x++
		case 'S':
			x--
		case 'E':
			y--
		case 'W':
			y++
		}
		// 把x和y 坐标单独计算，现加入向东走了2,向西走了4,那么 x的距离是2，如果改一下，把向东走改成向西走一步，那么距离就是5-1，可以看出，
		// 每改一次，增加的距离是2，但是最多能增加多少呢：最大增加到全部向西走：也就是i+1(i从0开始的，所以要加1)
		ans = max(ans, min(i+1, abs(x)+abs(y)+k*2))
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func maxDistance2(s string, k int) (ans int) {
	cnt := make(map[byte]int)

	f := func(a, b, left int) (int, int) {
		d := min(a, b, left)
		return abs(a-b) + d*2, d
	}

	for _, ch := range s {
		cnt[byte(ch)]++
		left := k
		x, d1 := f(cnt['N'], cnt['S'], left)
		left -= d1
		y, d2 := f(cnt['E'], cnt['W'], left)
		left -= d2
		ans = max(ans, x+y)
	}
	return
}
