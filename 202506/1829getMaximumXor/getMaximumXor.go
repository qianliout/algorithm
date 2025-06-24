package main

func main() {

}

// getMaximumXor 解决最大异或查询问题
// 核心思想：要使异或结果最大，k 应该是当前异或和的"补码"（在 maximumBit 位范围内）
func getMaximumXor(nums []int, maximumBit int) []int {
	// 第一步：计算整个数组的异或和
	// 异或运算的性质：a ^ b ^ b = a（自己和自己异或等于0）
	sumOr := 0
	for _, ch := range nums {
		sumOr = sumOr ^ ch
	}

	// 第二步：创建掩码 mask = 2^maximumBit - 1
	// 例如：maximumBit=3，则 mask = 2^3 - 1 = 8 - 1 = 7 = 111(二进制)
	// mask 的作用：表示在 maximumBit 位范围内的最大值，所有位都是1
	mask := 1<<maximumBit - 1

	n := len(nums)
	ans := make([]int, n)

	// 第三步：从后往前处理每个查询
	// 每次查询后要删除数组的最后一个元素
	for i := n - 1; i >= 0; i-- {
		// 关键理解：要使 (sumOr ^ k) 最大，k 应该是 sumOr 在 maximumBit 位内的"反码"
		//
		// 原理解释：
		// 1. 我们想要 sumOr ^ k 的结果尽可能大
		// 2. 在二进制中，数值越大意味着高位的1越多
		// 3. 对于 sumOr 的每一位：
		//    - 如果 sumOr 的第i位是0，我们希望 k 的第i位是1，这样 0^1=1
		//    - 如果 sumOr 的第i位是1，我们希望 k 的第i位是0，这样 1^0=1
		// 4. 这正好是 sumOr 的按位取反操作
		// 5. 但由于 k < 2^maximumBit，我们只关心低 maximumBit 位
		// 6. sumOr ^ mask 正好实现了在 maximumBit 位范围内的按位取反
		ans[n-1-i] = sumOr ^ mask

		// 第四步：模拟删除最后一个元素
		// 由于异或的性质：a ^ b ^ b = a
		// 如果 sumOr = nums[0] ^ nums[1] ^ ... ^ nums[i]
		// 那么 sumOr ^ nums[i] = nums[0] ^ nums[1] ^ ... ^ nums[i-1]
		// 这就相当于删除了 nums[i]
		sumOr = sumOr ^ nums[i]
	}

	return ans
}

/*
给你一个 有序 数组 nums ，它由 n 个非负整数组成，同时给你一个整数 maximumBit 。你需要执行以下查询 n 次：
找到一个非负整数 k < 2maximumBit ，使得 nums[0] XOR nums[1] XOR ... XOR nums[nums.length-1] XOR k 的结果 最大化 。k 是第 i 个查询的答案。
从当前数组 nums 删除 最后 一个元素。
请你返回一个数组 answer ，其中 answer[i]是第 i 个查询的结果。
*/
