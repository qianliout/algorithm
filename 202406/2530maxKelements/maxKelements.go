package main

import (
	"container/heap"

	. "outback/algorithm/common/commonHeap"
)

func main() {

}

// 直接模拟会timeout
func maxKelements1(nums []int, k int) int64 {
	ans := 0
	for k > 0 {
		idx := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] > nums[idx] {
				idx = i
			}
		}
		ans += nums[idx]
		nums[idx] = (nums[idx] + 2) / 3
		k--

	}
	return int64(ans)
}

func maxKelements(nums []int, k int) int64 {
	ans := 0
	mh := make(MaxHeap, 0)

	for _, ch := range nums {
		heap.Push(&mh, ch)
	}

	for k > 0 {
		top := heap.Pop(&mh).(int)
		ans += top
		heap.Push(&mh, (top+2)/3)
		k--
	}
	return int64(ans)
}
