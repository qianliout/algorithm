package main

import (
	"fmt"
)

func main() {
	fmt.Println(angleClock(12, 30))
}
func angleClock1(hour int, minutes int) float64 {
	m := float64(minutes) / float64(5)
	h := float64(hour%12) + float64(minutes)/float64(60)
	x := float64(360) * (abs(m-h) / float64(12))
	return min(x, 360-x)
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

/*
由于圆周有 360 度，时钟上有 12 个数字，因此一小时内时针顺时针旋转的度数是 30 度。为方便计算，将小时数 hour 限定在范围 [0,11] 中，如果 hour=12 则将 hour 更新为 0。根据 hour 的值得到原始角度是 30×hour 度。
由于一小时有 60 分钟，一小时内时针顺时针旋转 30 度，因此一分钟内时针顺时针旋转 0.5 度。根据 minutes 的值得到偏移角度是 0.5×minutes 度。
*/

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}
	h := float64(hour*30) + float64(minutes)*0.5
	m := float64(minutes * 6)
	x := abs(h - m)
	return min(x, 360-x)
}
