package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// findMedian 函数用于找到数组的中位数
// 中位数定义：
// - 如果数组长度为奇数，中位数是中间位置的数
// - 如果数组长度为偶数，中位数是中间两个数的平均值
func findMedian(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n&1 == 1 {
		return float64(quickSelect(nums, 0, n-1, n/2))
	}

	a := quickSelect(nums, 0, n-1, n/2)
	b := quickSelect(nums, 0, n-1, n/2-1)
	return float64(a+b) / 2

}

func quickSelect(nums []int, left, right int, k int) int {
	if left >= right {
		return nums[left]
	}
	p := partition(nums, left, right)
	if p == k {
		return nums[p]
	}
	if k < p {
		return quickSelect(nums, left, p-1, k)
	}
	return quickSelect(nums, p+1, right, k)
}

func partition(nums []int, left, right int) int {
	if left >= right {
		return left
	}
	// 可以随机选一个，这里就直接选末尾的
	pivot := nums[right]
	le := left
	ri := right - 1
	// 这里一定是le<=ri ，为啥呢,因为le表示是小堆写入的位置，也就是现在有一个数要写入小堆时的写入位置
	// 当le==ri时，虽然不会改变nums[le]可nums[ri]的值，但是到底是le++,还是ri--还不确定
	for le <= ri {
		if nums[le] <= pivot {
			le++
		} else {
			nums[le], nums[ri] = nums[ri], nums[le]
			ri--
		}
	}
	nums[right], nums[le] = nums[le], nums[right]
	return le
}

// partition 函数将数组按照pivot进行分区
// 分区后，pivot左边的元素都小于pivot，右边的元素都大于等于pivot
// 返回pivot的最终位置
func partition2(nums []int, left, right int) int {
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

func findMedian2(nums []int) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sort.Ints(nums)

	if n%2 == 1 {
		return float64(nums[n/2])
	}
	return float64(nums[n/2]+nums[n/2-1]) / 2

}

func main() {
	// 设置随机数种子，确保每次运行结果不同
	rand.Seed(time.Now().UnixNano())

	// 创建测试数据：1000个随机浮点数
	data := make([]int, 1e3+999) // 1000个元素
	for i := range data {
		data[i] = int(rand.Float64() * 1000)
	}

	// 计算并输出中位数
	median := findMedian(data)
	fmt.Println("Median:", median, findMedian2(data))
}
