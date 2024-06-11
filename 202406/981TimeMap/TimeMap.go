package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(sort.Search(len(nums), func(k int) bool { return nums[k] > 2 }))
	fmt.Println(sort.SearchInts(nums, 0))
}

type TimeMap struct {
	Data  map[string][]string
	Index map[string][]int // 对应的key 的 时间
}

func Constructor() TimeMap {
	return TimeMap{
		Data:  make(map[string][]string),
		Index: make(map[string][]int),
	}

}

func (this *TimeMap) Set(key string, value string, t int) {
	this.Data[key] = append(this.Data[key], value)
	this.Index[key] = append(this.Index[key], t)
}

// func (this *TimeMap) Get(key string, t int) string {
// 	nums := this.Index[key]
// 	if len(nums) == 0 {
// 		return ""
// 	}
// 	data := this.Data[key]
//
// 	idx := sort.Search(len(nums), func(k int) bool { return nums[k] <= t })
// 	if idx >= len(nums) {
// 		return data[len(data)-1]
// 	}
// 	return data[idx]
// }

func (this *TimeMap) Get(key string, t int) string {
	nums := this.Index[key]
	if len(nums) == 0 {
		return ""
	}
	data := this.Data[key]
	// idx==0 说明 t 比nums[0]都小
	idx := sort.Search(len(nums), func(k int) bool { return nums[k] > t })
	if idx <= 0 {
		return ""
	}
	return data[idx-1]
}
