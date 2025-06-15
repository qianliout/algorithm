package main

import "fmt"

func main() {
	// 测试原始错误的代码
	fmt.Println("=== 原始错误代码测试 ===")
	nums1 := []int{0, 1, 0, 3, 12}
	fmt.Printf("原数组: %v\n", nums1)
	moveZeroesWrong(nums1)
	fmt.Printf("错误结果: %v\n", nums1)
	fmt.Println()

	// 测试正确的代码
	fmt.Println("=== 正确代码测试 ===")
	testCases := [][]int{
		{0, 1, 0, 3, 12},
		{0, 0, 1},
		{1, 2, 3},
		{0, 0, 0},
		{1},
	}

	for i, nums := range testCases {
		original := make([]int, len(nums))
		copy(original, nums)
		fmt.Printf("测试%d - 原数组: %v\n", i+1, original)
		moveZeroes(nums)
		fmt.Printf("测试%d - 结果:   %v\n", i+1, nums)
		fmt.Println()
	}

	// 可视化执行过程
	fmt.Println("=== 执行过程可视化 ===")
	visualNums := []int{0, 1, 0, 3, 12}
	visualizeMoveZeroes(visualNums)
}

// 你的原始错误代码（重命名以便对比）
func moveZeroesWrong(nums []int) {
	n := len(nums)
	i := 0
	pre := 0
	for i < n && pre < n {
		if nums[i] == 0 {
			i++
			continue
		}
		if nums[i] != 0 {
			nums[i], nums[pre] = nums[pre], nums[i]
			pre++
		}
		if nums[i] == 0 {
			i++
		}
	}
}

// ✅ 正确解法1：双指针法（推荐）
func moveZeroes1(nums []int) {
	/*
		核心思想：使用双指针
		- left: 指向下一个非零元素应该放置的位置
		- right: 遍历数组，寻找非零元素

		算法步骤：
		1. left 指针指向第一个位置
		2. right 指针遍历数组
		3. 当 right 指向非零元素时，将其与 left 位置交换
		4. left 指针前移
	*/

	left := 0 // 指向下一个非零元素应该放置的位置

	// right 指针遍历整个数组
	for right := 0; right < len(nums); right++ {
		// 如果当前元素非零
		if nums[right] != 0 {
			// 交换元素（即使 left == right 也没关系）
			nums[left], nums[right] = nums[right], nums[left]
			// left 指针前移
			left++
		}
		// right 指针自动前移（for循环）
	}
}

// ✅ 正确解法2：优化版双指针（避免不必要的交换）
func moveZeroesOptimized(nums []int) {
	/*
		优化思想：只有当 left != right 时才交换
		这样可以避免元素与自己交换的无意义操作
	*/

	left := 0 // 指向下一个非零元素应该放置的位置

	// 第一遍：将所有非零元素移到前面
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			if left != right { // 只有位置不同时才交换
				nums[left] = nums[right]
			}
			left++
		}
	}

	// 第二遍：将剩余位置填充为0
	for left < len(nums) {
		nums[left] = 0
		left++
	}
}

// 📊 执行过程可视化
func visualizeMoveZeroes(nums []int) {
	fmt.Printf("原数组: %v\n", nums)

	left := 0
	for right := 0; right < len(nums); right++ {
		fmt.Printf("步骤%d: left=%d, right=%d, nums[right]=%d\n",
			right+1, left, right, nums[right])

		if nums[right] != 0 {
			fmt.Printf("  -> 非零元素，交换 nums[%d] 和 nums[%d]\n", left, right)
			nums[left], nums[right] = nums[right], nums[left]
			fmt.Printf("  -> 数组变为: %v\n", nums)
			left++
		} else {
			fmt.Printf("  -> 零元素，跳过\n")
		}
		fmt.Printf("  -> 当前数组: %v\n", nums)
		fmt.Println()
	}

	fmt.Printf("最终结果: %v\n", nums)
}

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func moveZeroes(nums []int) {
	n := len(nums)
	left := 0
	for right := 0; right < n; right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
