package main

func main() {

}

/*
给你一个下标从 0 开始的整数数组 nums ，其长度是 2 的幂。

对 nums 执行下述算法：

	设 n 等于 nums 的长度，如果 n == 1 ，终止 算法过程。否则，创建 一个新的整数数组 newNums ，新数组长度为 n / 2 ，下标从 0 开始。
	对于满足 0 <= i < n / 2 的每个 偶数 下标 i ，将 newNums[i] 赋值 为 min(nums[2 * i], nums[2 * i + 1]) 。
	对于满足 0 <= i < n / 2 的每个 奇数 下标 i ，将 newNums[i] 赋值 为 max(nums[2 * i], nums[2 * i + 1]) 。
	用 newNums 替换 nums 。
	从步骤 1 开始 重复 整个过程。
*/
func minMaxGame(nums []int) int {
	arr := dfs(nums)
	return arr[0]
}

func dfs(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	n := len(nums)
	arr := make([]int, n/2)
	for i := range arr {
		if i&1 == 0 {
			arr[i] = min(nums[2*i], nums[2*i+1])
		} else {
			arr[i] = max(nums[2*i], nums[2*i+1])
		}
	}
	return dfs(arr)
}
