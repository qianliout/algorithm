package main

func main() {

}

/*
给你一个数组 nums 和一个整数 target 。
请你返回 非空不重叠 子数组的最大数目，且每个子数组中数字和都为 target 。
*/

func maxNonOverlapping(nums []int, target int) int {
	n := len(nums)
	sum := 0
	idx := make(map[int]int)
	idx[0] = -1 // 表示空前缀的前缀和为0，索引为-1
	ans := 0
	last := -1
	// 使用前缀和 + 哈希表的贪心策略，一旦找到和为 target 的子数组就立即选择它。
	// 如果我们延迟选择某个子数组，等待更好的组合，实际上不会获得更多的子数组数量
	// 早选择可以为后续留出更多空间，不会影响最优解
	for i := 0; i < n; i++ {
		sum += nums[i]
		pre, ok := idx[sum-target]
		// 找的是前缀和，前缀和是开区间,所以区区间内的第一个元素是 pre+1,又因为不能是空集，所以是 pre+1>last
		// pre 是前缀和 sum - target 对应的索引
		// pre+1 是当前子数组的起始位置
		// last 是上一个选中子数组的结束位置
		// pre+1 > last 确保当前子数组在上一个子数组之后开始
		if ok && pre+1 > last {
			ans++
			last = i
		}
		idx[sum] = i // 不提前算好，有两个原因，1是，因为有负数，会有多个相同的前缀和，2，没有办法判断是否重叠
	}
	return ans
}
