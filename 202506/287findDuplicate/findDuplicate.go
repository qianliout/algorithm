package main

import (
	"fmt"
)

func main() {
	// 测试用例1：[1,3,4,2,2]
	nums1 := []int{1, 3, 4, 2, 2}
	fmt.Printf("数组: %v\n", nums1)
	fmt.Printf("重复数字: %d\n", findDuplicate(nums1))
	visualizeArray(nums1)
	fmt.Println()

	// 测试用例2：[3,3,3,3,3]
	nums2 := []int{3, 3, 3, 3, 3}
	fmt.Printf("数组: %v\n", nums2)
	fmt.Printf("重复数字: %d\n", findDuplicate(nums2))
	visualizeArray(nums2)
}

// 可视化数组的链表结构
func visualizeArray(nums []int) {
	fmt.Println("链表可视化:")
	fmt.Print("索引: ")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Print("值:   ")
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Print("链表: ")
	visited := make(map[int]bool)
	current := 0

	for !visited[current] {
		fmt.Printf("%d", current)
		visited[current] = true
		next := nums[current]
		if visited[next] {
			fmt.Printf(" -> %d (环的入口)", next)
			break
		} else {
			fmt.Print(" -> ")
		}
		current = next
	}
	fmt.Println()
}

func findDuplicate(nums []int) int {
	/*
		核心思想：将数组看作链表，使用弗洛伊德环检测算法

		数组到链表的映射：
		- 索引 i 看作链表节点
		- nums[i] 看作从节点 i 指向节点 nums[i] 的指针
		- 由于存在重复数字，必然形成环

		例如：nums = [1,3,4,2,2]
		索引:  0 1 2 3 4
		值:    1 3 4 2 2

		链表关系：
		0 -> 1 -> 3 -> 2 -> 4 -> 2 (形成环，2是环的入口)
	*/

	// 阶段1：使用快慢指针检测环的存在
	// slow 每次走一步，fast 每次走两步
	slow, fast := nums[0], nums[nums[0]]

	// 快慢指针在环内相遇
	for slow != fast {
		slow = nums[slow]       // 慢指针走一步
		fast = nums[nums[fast]] // 快指针走两步
	}

	/*
		阶段2：找到环的入口（即重复的数字）

		数学原理：
		设链表头到环入口距离为 a，环的长度为 b
		当快慢指针相遇时：
		- 慢指针走了 a + x 步（x 是在环内走的距离）
		- 快指针走了 a + x + kb 步（k 是快指针多走的圈数）
		- 由于快指针速度是慢指针的2倍：2(a + x) = a + x + kb
		- 化简得：a + x = kb，即 a = kb - x

		这意味着：从链表头走 a 步 = 从相遇点走 kb-x 步
		由于 kb-x = (k-1)b + (b-x)，相当于从相遇点走 b-x 步到环入口
	*/

	// 将 fast 重置到链表头（索引0）
	fast = 0

	// slow 从相遇点开始，fast 从头开始，都每次走一步
	// 它们会在环的入口相遇
	for fast != slow {
		slow = nums[slow] // 从相遇点继续走
		fast = nums[fast] // 从头开始走
	}

	// 相遇点就是环的入口，即重复的数字
	return slow
}

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
