package main

import "sort"

func main() {

}

func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	ids := make([]int, n)
	for i := range nums {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool { return nums[ids[i]] < nums[ids[j]] })
	ans := n // 初值不是0,每个元素都要删除一次
	for i := 1; i < n; i++ {
		if ids[i] < ids[i-1] {
			ans += n - i
		}
	}
	return int64(ans)
}
