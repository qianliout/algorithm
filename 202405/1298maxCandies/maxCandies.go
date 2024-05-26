package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCandies([]int{1, 0, 1, 0}, []int{7, 5, 4, 100},
		[][]int{{}, {}, {1}, {}}, [][]int{{1, 2}, {3}, {}, {}},
		[]int{0}))
}

type Box struct {
	Idx    int
	HasKey bool
	HasBox bool
}

func (vi Box) open() bool {
	return vi.HasBox && vi.HasKey
}
func empty(aa []Box) bool {
	for _, a := range aa {
		if a.open() {
			return true
		}
	}
	return false
}

func merge(aa []Box, status []int) []Box {
	ans := make([]Box, 0)
	exi := make(map[int]*Box)
	for i := range aa {
		b := aa[i]
		if _, ok := exi[b.Idx]; !ok {
			exi[b.Idx] = &b
			continue
		}
		if b.HasBox {
			exi[b.Idx].HasBox = true
		}
		if b.HasKey {
			exi[b.Idx].HasKey = true
		}
	}
	for k := range exi {
		v := exi[k]
		ans = append(ans, *v)
	}
	for i := range ans {
		if ans[i].HasBox && status[ans[i].Idx] == 1 {
			ans[i].HasKey = true
		}
	}
	return ans
}

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {

	queue := make([]Box, 0)
	for _, b := range initialBoxes {
		queue = append(queue, Box{
			Idx:    b,
			HasKey: status[b] == 1,
			HasBox: true,
		})
	}
	ans := 0
	visit := make([]bool, len(status))
	for len(queue) > 0 && empty(queue) {
		lev := make([]Box, 0)
		for _, no := range queue {
			// 没有能找打开，那就去下一轮测试
			if !no.open() {
				lev = append(lev, no)
				continue
			}
			if visit[no.Idx] {
				continue
			}
			ans += candies[no.Idx]
			visit[no.Idx] = true
			// 找有钥匙的
			for _, k := range keys[no.Idx] {
				lev = append(lev, Box{
					Idx:    k,
					HasKey: true,
					HasBox: false,
				})
			}
			for _, b := range containedBoxes[no.Idx] {
				lev = append(lev, Box{
					Idx:    b,
					HasKey: false,
					HasBox: true,
				})
			}
		}

		lev = merge(lev, status)
		queue = lev
	}
	return ans
}
