package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(smallestRange([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}))
	// fmt.Println(smallestRange([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}))
}

type pair struct {
	id1, id2, value int // id1表示在 nums 中的下标，id2表示 在nums[id1]的下标
}

func smallestRange(nums [][]int) []int {
	n := len(nums)
	data := make([]pair, 0)
	for i := range nums {
		for j, ch := range nums[i] {
			data = append(data, pair{id1: i, id2: j, value: ch})
		}
	}
	sort.Slice(data, func(i, j int) bool {
		if data[i].value < data[j].value {
			return true
		} else if data[i].value == data[j].value {
			return data[i].id1 < data[j].id1
		} else if data[i].value > data[j].value {
			return false
		}
		return true
	})
	// for i := range data {
	// 	fmt.Println(data[i].id1)
	// }

	// 滑动窗口
	le, ri := 0, 0
	wind := make([]pair, 0)
	ans := make([]int, 0)
	used := make(map[int]int)
	minSub := math.MaxInt
	for le <= ri && ri < len(data) {
		wind = append(wind, data[ri])
		used[data[ri].id1]++
		for len(used) >= n {
			su := wind[ri].value - wind[le].value
			if su < minSub {
				minSub = su
				ans = []int{wind[le].value, wind[ri].value}
			}
			used[data[le].id1]--
			if used[data[le].id1] == 0 {
				delete(used, data[le].id1)
			}
			le++
		}
		ri++
	}
	return ans
}

// 按组进行滑窗，保证一个窗口的组满足kkk组后在记录窗口的最小区间值
func need(wind []pair, used map[int]int, n int) bool {
	u := 0
	for _, v := range used {
		if v > 0 {
			u++
		}
	}
	return u >= n
}
