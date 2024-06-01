package main

func main() {

}

// https://leetcode.cn/problems/global-and-local-inversions/solutions/1973365/by-ac_oier-jc7a/
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 局部倒置是全局倒置的子集，又因为数组范围在 0---->n-1
		if abs(nums[i]-i) > 1 {
			return false
		}
	}
	return true
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
