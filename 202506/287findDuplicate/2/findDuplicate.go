package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

/*
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
*/

/*

2. 为什么不能从 0 开始？
如果数组元素范围是 [0, n-1]，会出现以下问题：

问题1：无法形成有效的环

如果 nums[i] = 0，那么 nums[0] 就是环的起点
但索引 0 本身就是起点，这会导致算法逻辑混乱
问题2：边界访问问题

当 nums[i] = 0 时，下一步访问 nums[0]
如果数组长度是 n+1，索引范围是 [0, n]，那么值域 [0, n-1] 无法覆盖所有可能的索引
3. 范围 [1, n] 的优势
完美的映射关系：

数组长度：n+1（索引 0 到 n）
元素值域：[1, n]（正好对应索引 1 到 n）
索引 0 作为虚拟起点，不会被任何元素值指向
*/
