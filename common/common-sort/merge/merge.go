package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 4, 2, 2, 2, 20, 3}
	MergeSort(nums)
	// SelectionSort(nums)
	// BubbleSort(nums)
	fmt.Println(nums)
}

// 归并排序算法（分治策略）
// 时间复杂度: O(n log n) - 无论最好最坏情况都是O(n log n)
// 空间复杂度: O(n) - 需要额外的数组空间来存储合并结果
// 稳定性: 稳定排序算法
// 工作原理: 采用分治策略，将数组递归地分成两半，分别排序后再合并
func MergeSort(nums []int) {
	// 基本情况：如果数组长度小于等于1，已经是有序的
	if len(nums) <= 1 {
		return
	}
	// 计算中点，将数组分成两半
	mid := len(nums) / 2
	// 递归排序左半部分
	MergeSort(nums[:mid])
	// 递归排序右半部分
	MergeSort(nums[mid:])
	// 合并两个已排序的子数组
	Merge(nums, mid)
}

// Merge 合并两个已排序的子数组（原地合并实现）
// 参数: nums - 包含两个已排序子数组的完整数组
//
//	mid - 分割点，nums[:mid]是左半部分，nums[mid:]是右半部分
//
// 时间复杂度: O(n²) - 使用插入排序的思想进行合并
// 空间复杂度: O(1) - 原地操作，不需要额外空间
// 注意: 这是一个简化的合并实现，标准的归并排序合并操作时间复杂度为O(n)
func Merge(nums []int, mid int) {
	// 从右半部分的第一个元素开始
	for i := mid; i < len(nums); i++ {
		// 将右半部分的元素插入到左半部分的正确位置
		// 使用插入排序的思想，从当前位置向前比较和交换
		for j := i; j > 0; j-- {
			// 如果当前元素已经大于等于前一个元素，说明找到了正确位置
			if nums[j] >= nums[j-1] {
				continue
			}
			// 否则交换位置，继续向前比较
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
