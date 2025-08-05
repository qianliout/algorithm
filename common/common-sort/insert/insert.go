package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 4, 2, 2, 2, 20, 3}
	// insertSort(nums)
	// SelectionSort(nums)
	// BubbleSort(nums)
	ShellSort(nums)
	fmt.Println(nums)

}

// 插入排序算法实现1
// 时间复杂度: O(n²) - 最坏情况下需要O(n²)次比较和交换
// 空间复杂度: O(1) - 只需要常数级别的额外空间
// 稳定性: 稳定排序算法
// 工作原理: 将数组分为已排序和未排序两部分，每次从未排序部分取出一个元素插入到已排序部分的正确位置
func insertSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 从第二个元素开始，因为第一个元素可以认为是已排序的
	for i := 1; i < len(nums); i++ {
		// 将当前元素与已排序部分的元素从后往前比较
		// 如果当前元素小于前一个元素，则交换位置
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// 插入排序算法实现2（优化版本）
// 时间复杂度: O(n²) - 最坏情况下需要O(n²)次比较和交换
// 空间复杂度: O(1) - 只需要常数级别的额外空间
// 稳定性: 稳定排序算法
// 优化点: 当找到正确位置时提前退出内层循环，减少不必要的比较
func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 将当前元素与已排序部分的元素从后往前比较
		for j := i; j > 0; j-- {
			// 如果当前元素已经大于等于前一个元素，说明找到了正确位置，可以退出
			if nums[j] >= nums[j-1] {
				break
			}
			// 否则交换位置，继续向前比较
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// 选择排序算法
// 时间复杂度: O(n²) - 无论最好最坏情况都需要O(n²)次比较
// 空间复杂度: O(1) - 只需要常数级别的额外空间
// 稳定性: 不稳定排序算法（相同元素的相对位置可能改变）
// 工作原理: 将数组分为已排序和未排序两部分，每次从未排序部分选择最小的元素放到已排序部分的末尾
func SelectionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 遍历数组，i表示已排序部分的末尾位置
	for i := 0; i < len(nums); i++ {
		// 找到未排序部分中最小元素的索引
		mia := i
		for j := i + 1; j < len(nums); j++ {
			if nums[mia] > nums[j] {
				mia = j
			}
		}
		// 将找到的最小元素与当前位置交换
		nums[i], nums[mia] = nums[mia], nums[i]
	}
}

// 冒泡排序算法
// 时间复杂度: O(n²) - 最坏情况下需要O(n²)次比较和交换
// 空间复杂度: O(1) - 只需要常数级别的额外空间
// 稳定性: 稳定排序算法
// 工作原理: 重复遍历数组，每次比较相邻的两个元素，如果顺序错误则交换它们
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 外层循环控制排序轮数
	for i := 0; i < len(nums); i++ {
		// 内层循环进行相邻元素比较和交换
		// 注意：这里的内层循环逻辑有问题，应该是冒泡的标准实现
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// 希尔排序算法（插入排序的改进版本）
// 时间复杂度: O(n^1.3) - 平均情况下比插入排序更快
// 空间复杂度: O(1) - 只需要常数级别的额外空间
// 稳定性: 不稳定排序算法
// 工作原理: 使用不同的间隔序列对数组进行分组，对每组使用插入排序，逐渐减小间隔直到为1
func ShellSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 使用希尔增量序列，初始间隔为数组长度的一半
	for gap := len(nums) / 2; gap > 0; gap = gap / 2 {
		// 对每个间隔进行插入排序
		for i := gap; i < len(nums); i++ {
			// 保存当前要插入的元素
			tem := nums[i]
			j := i
			// 在间隔为gap的序列中进行插入排序
			for j := i; j >= gap; j = j - gap {
				// 如果前一个元素已经小于等于当前元素，找到插入位置
				if nums[j-gap] >= tem {
					break
				}
				// 否则将前一个元素向后移动
				nums[j] = nums[j-gap]
			}
			// 将保存的元素插入到正确位置
			nums[j] = tem
		}
	}
}
