package main

import (
	"sort"
)

func main() {

}

type CountIntervals struct {
	Data []pair
	Cnt  int
}

func Constructor() CountIntervals {
	return CountIntervals{
		Data: make([]pair, 0),
		Cnt:  0,
	}
}

func (this *CountIntervals) Add(left int, right int) {
	this.Data = append(this.Data, pair{left, right})

	pairs, cnt := merge(this.Data)
	this.Data = pairs
	this.Cnt = cnt

}

func (this *CountIntervals) Count() int {
	return this.Cnt
}

type pair struct {
	start, end int
}

// 这种方式可以完成，但是会超时
func merge(data []pair) ([]pair, int) {
	cout := 0
	sort.Slice(data, func(i, j int) bool {
		if data[i].start < data[j].start {
			return true
		} else if data[i].start > data[j].start {
			return false
		} else {
			return data[i].end < data[j].end
		}
	})
	ans := make([]pair, 0)
	start, end := data[0].start, data[0].end
	for i := 1; i < len(data); i++ {
		if data[i].start > end {
			ans = append(ans, pair{start: start, end: end})
			cout += end - start + 1
			start, end = data[i].start, data[i].end
		} else {
			end = max(end, data[i].end)
		}
	}
	ans = append(ans, pair{start: start, end: end})
	cout += end - start + 1
	return ans, cout
}
