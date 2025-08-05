package main

import (
	"fmt"
	"math/rand"
	"time"
)

// findMedian 函数用于找到数组的中位数
// 中位数定义：
// - 如果数组长度为奇数，中位数是中间位置的数
// - 如果数组长度为偶数，中位数是中间两个数的平均值
func findMedian(nums []float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 如果是奇数长度，返回中间那个数
	// 例如：[1,2,3,4,5] -> 中位数是第3个数（索引2）
	if n%2 == 1 {
		return quickSelect(nums, 0, n-1, n/2)
	}
	// 如果是偶数长度，返回中间两个数的平均值
	// 例如：[1,2,3,4] -> 中位数是(第2个数 + 第3个数) / 2
	return 0.5 * (quickSelect(nums, 0, n-1, n/2-1) + quickSelect(nums, 0, n-1, n/2))
}

// quickSelect 函数使用快速选择算法找到数组中第k小的元素
// 这是快速排序的变种，时间复杂度平均为O(n)，最坏情况为O(n²)
// 参数说明：
// - nums: 输入数组
// - left, right: 当前搜索范围的左右边界
// - k: 要找到的第k小的元素（从0开始计数）
func quickSelect(nums []float64, left, right, k int) float64 {
	// 如果搜索范围只有一个元素，直接返回
	if left == right {
		return nums[left]
	}

	// 使用partition函数将数组分为两部分，返回pivot的最终位置
	pivotIndex := partition(nums, left, right)

	// 如果pivot的位置正好是我们要找的第k个位置
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		// 如果k在pivot左边，在左半部分继续查找
		return quickSelect(nums, left, pivotIndex-1, k)
	}
	// 如果k在pivot右边，在右半部分继续查找
	return quickSelect(nums, pivotIndex+1, right, k)
}

// partition 函数将数组按照pivot进行分区
// 分区后，pivot左边的元素都小于pivot，右边的元素都大于等于pivot
// 返回pivot的最终位置
func partition(nums []float64, left, right int) int {
	// 随机选择pivot，避免最坏情况（已排序数组）
	pivotIndex := rand.Intn(right-left+1) + left
	pivot := nums[pivotIndex]

	// 将pivot移到数组末尾，方便后续处理
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// storeIndex指向小于pivot的元素应该放置的位置
	storeIndex := left

	// 遍历数组（除了pivot），将小于pivot的元素移到左边
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			// 交换元素，将小于pivot的元素移到storeIndex位置
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}

	// 将pivot移到最终位置（storeIndex）
	nums[right], nums[storeIndex] = nums[storeIndex], nums[right]
	return storeIndex
}

func main() {
	// 设置随机数种子，确保每次运行结果不同
	rand.Seed(time.Now().UnixNano())

	// 创建测试数据：1000个随机浮点数
	data := make([]float64, 1e3) // 1000个元素
	for i := range data {
		data[i] = rand.Float64() * 1000
	}

	// 计算并输出中位数
	median := findMedian(data)
	fmt.Printf("Median: %f\n", median)
}
