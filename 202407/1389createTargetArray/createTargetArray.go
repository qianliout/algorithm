package main

import (
	"fmt"
)

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}

// 一定要深赋值，不然就会出错
func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		nu, id := nums[i], index[i]

		pre := append([]int{}, ans[:id]...)
		after := append([]int{}, ans[id:]...)

		pre = append(pre, nu)
		pre = append(pre, after...)
		ans = append([]int{}, pre...)
	}
	return ans
}
