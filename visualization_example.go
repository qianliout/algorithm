package main

import "fmt"

func visualizeLinkedList(nums []int, title string) {
	fmt.Printf("\n=== %s ===\n", title)
	fmt.Printf("数组: %v\n", nums)
	fmt.Printf("索引: ")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n值:   ")
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n\n链表路径追踪:\n")

	visited := make(map[int]bool)
	current := 0
	step := 0

	for step < 10 { // 最多追踪10步
		if visited[current] {
			fmt.Printf("步骤 %d: 到达索引 %d (已访问过) - 检测到环!\n", step, current)
			break
		}

		visited[current] = true
		fmt.Printf("步骤 %d: 索引 %d -> ", step, current)

		if current >= len(nums) {
			fmt.Printf("越界错误!\n")
			break
		}

		next := nums[current]
		fmt.Printf("值 %d (下一个索引)\n", next)

		if next >= len(nums) || next < 0 {
			fmt.Printf("下一个索引 %d 越界!\n", next)
			break
		}

		current = next
		step++
	}
}

func main() {
	// 标准情况：[1, n] 范围
	fmt.Println("标准 findDuplicate 问题设计:")
	fmt.Println("- 数组长度: n+1")
	fmt.Println("- 元素范围: [1, n]")
	fmt.Println("- 重复元素: 恰好1个")

	nums1 := []int{1, 3, 4, 2, 2} // 长度5，范围[1,4]
	visualizeLinkedList(nums1, "标准情况: [1, n] 范围")

	// 问题情况1：[0, n-1] 范围
	fmt.Println("\n" + "="*50)
	fmt.Println("如果改为 [0, n-1] 范围会出现的问题:")

	nums2 := []int{0, 2, 3, 1, 2} // 长度5，范围[0,3]，但索引4无法被访问
	visualizeLinkedList(nums2, "问题情况1: [0, n-1] 范围 - 包含0")

	// 问题情况2：[0, n-1] 范围，无0
	nums3 := []int{1, 2, 3, 1, 2} // 长度5，范围[1,3]，但缺少对索引4的引用
	visualizeLinkedList(nums3, "问题情况2: [0, n-1] 范围 - 不包含0但范围不匹配")

	// 解释核心原理
	fmt.Println("\n" + "="*50)
	fmt.Println("核心原理解释:")
	fmt.Println("1. Floyd环检测需要将数组转换为链表")
	fmt.Println("2. 数组索引 = 链表节点")
	fmt.Println("3. 数组值 = 指向下一个节点的指针")
	fmt.Println("4. 要形成有效环，需要:")
	fmt.Println("   - 所有值都是有效的数组索引")
	fmt.Println("   - 有一个固定起点(索引0)")
	fmt.Println("   - 值域[1,n]确保不会指回起点0")
	fmt.Println("   - 长度n+1确保有重复元素")

	fmt.Println("\n为什么 [1, n] 是完美设计:")
	fmt.Println("✓ 数组长度 n+1，索引范围 [0, n]")
	fmt.Println("✓ 元素值域 [1, n]，正好覆盖索引 [1, n]")
	fmt.Println("✓ 索引0作为起点，永远不会被指向")
	fmt.Println("✓ 由鸽巢原理，n+1个元素放入n个值，必有重复")
	fmt.Println("✓ 重复元素形成环，Floyd算法可以检测")
}
