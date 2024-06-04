package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
}

/*
我们定义两个变量 firstfirstfirst 和 lastlastlast 分别表示第一个人和最后一个人的位置，用变量 ddd 表示两个人之间的最大距离。
然后遍历数组 seatsseatsseats，如果当前位置有人，如果此前 lastlastlast 更新过，说明此前有人，此时更新 d=max⁡(d,i−last)d = \max(d, i - last)d=max(d,i−last)；如果此前 firstfirstfirst 没有更新过，说明此前没有人，此时更新 first=ifirst = ifirst=i。接下来更新 last=ilast = ilast=i。
最后返回 max⁡(first,n−last−1,d/2)\max(first, n - last - 1, d / 2)max(first,n−last−1,d/2) 即可。
*/
func maxDistToClosest(seats []int) int {
	inf := math.MinInt32 / 2
	d := 0
	fir, sec := inf, inf
	for i, ch := range seats {
		if ch > 0 {
			if sec != inf {
				d = max(d, i-sec)
			}
			if fir != inf {
				fir = i
			}
			sec = i
		}
	}
	return max(fir, len(seats)-sec-1, d/2)
}
