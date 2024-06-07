package main

func main() {

}

/*
你的任务是从范围 [0, n - 1] 内找出  2 个满足下述所有条件的下标 i 和 j ：

    abs(i - j) >= indexDifference 且
    abs(nums[i] - nums[j]) >= valueDifference

返回整数数组 answer。如果存在满足题目要求的两个下标，则 answer = [i, j] ；否则，answer = [-1, -1] 。如果存在多组可供选择的下标对，只需要返回其中任意一组即可。

注意：i 和 j 可能 相等 。
*/
// 数据量小，直接模拟
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	n := len(nums)
	ans := make([]int, 0)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if abs(j-i) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
				ans = append(ans, i, j)
				return ans
			}
		}
	}
	if len(ans) > 0 {
		return ans
	}
	return []int{-1, -1}
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
