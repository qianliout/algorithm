package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1}))
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1}))
}

var count int

func reversePairs(nums []int) int {
	count = 0
	mergeSort(nums, 0, len(nums)-1)
	return count
}

func merge(nums []int, l, mid, r int) {
	// 这里左边及右边数组的取数范围是一个难点，也是容易出错的点
	left := append([]int{}, nums[l:mid+1]...)
	right := append([]int{}, nums[mid+1:r+1]...)

	// 这里可以任务左边和右边分别都是有序数组了
	for i, j := 0, 0; i < len(left); i++ {
		for j < len(right) && left[i] > right[j]*2 {
			j++
		}
		count += j
	}

	// 技巧，下面在赋值时可以不用管边界
	left = append(left, math.MaxInt)
	right = append(right, math.MaxInt)
	i, j := 0, 0
	for k := l; k <= r; k++ {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
	}
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r + l) / 2
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}
