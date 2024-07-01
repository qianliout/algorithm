package main

import (
	"math"
)

func main() {

}

func sampleStats(count []int) []float64 {
	inf := math.MaxInt / 100
	mx := -inf
	mi := inf
	all := 0  // 值总和
	cnt := 0  // 个数总和
	mode := 0 // 从数

	for i, x := range count {
		if x > 0 {
			mx = max(mx, i)
			mi = min(mi, i)
			all += x * i
			cnt += x
			if count[mode] < x {
				mode = i
			}
		}
	}
	var medan float64
	if cnt&1 == 1 {
		medan = find(count, cnt/2+1)
	} else {
		a := find(count, cnt/2)
		b := find(count, cnt/2+1)
		medan = (a + b) / float64(2)
	}
	return []float64{float64(mi), float64(mx), float64(all) / float64(cnt), medan, float64(mode)}

}

// 找第 k 个元素
func find(count []int, k int) float64 {
	t := 0
	for i, x := range count {
		if x == 0 {
			continue
		}
		t += x
		if t >= k {
			return float64(i)
		}
	}
	return 0
}
