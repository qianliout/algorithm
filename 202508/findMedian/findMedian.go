package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findMedian(nums []float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 如果是奇数长度，返回中间那个数
	if n%2 == 1 {
		return quickSelect(nums, 0, n-1, n/2)
	}
	// 如果是偶数长度，返回中间两个数的平均值
	return 0.5 * (quickSelect(nums, 0, n-1, n/2-1) + quickSelect(nums, 0, n-1, n/2))
}

func quickSelect(nums []float64, left, right, k int) float64 {
	if left == right {
		return nums[left]
	}

	pivotIndex := partition(nums, left, right)
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return quickSelect(nums, left, pivotIndex-1, k)
	}
	return quickSelect(nums, pivotIndex+1, right, k)
}

func partition(nums []float64, left, right int) int {
	pivotIndex := rand.Intn(right-left+1) + left
	pivot := nums[pivotIndex]
	// 将pivot移到末尾
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	
	storeIndex := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}
	// 将pivot移到最终位置
	nums[right], nums[storeIndex] = nums[storeIndex], nums[right]
	return storeIndex
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// 示例数据
	data := make([]float64, 1e6) // 100万个元素
	for i := range data {
		data[i] = rand.Float64() * 1000
	}
	
	median := findMedian(data)
	fmt.Printf("Median: %f\n", median)
}