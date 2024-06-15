package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkArray([]int{60, 72, 87, 89, 63, 52, 64, 62, 31, 37, 57, 83, 98, 94, 92, 77, 94, 91, 87, 100, 91, 91, 50, 26}, 4))
	// fmt.Println(checkArray([]int{91, 91, 50, 26}, 4))
	// fmt.Println(checkArray([]int{1, 3, 1, 1}, 2))
	// fmt.Println(checkArray([]int{0, 45, 82, 98, 99}, 4))
}

func checkArray1(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n)
	d[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		d[i] = nums[i] - nums[i-1]
	}
	for i := 0; i < n; i++ {
		if d[i] == 0 {
			continue
		}
		if d[i] < 0 {
			return false
		}
		if d[i] > 0 {
			if i+k < n {
				d[i+k] += d[i]
				// d[i] = 0
			}
		}
	}
	// 看前面的更新是否已经把最后这k-1 个数都规0了
	// 最后k-1个不能操作的节点
	// 为啥是n -k +1
	for i := n - k + 1; i < n; i++ {
		if d[i] != 0 {
			return false
		}
	}
	return true
}

func checkArray(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	sumD := 0
	for i, x := range nums {
		sumD += d[i]
		x += sumD
		if x == 0 { // 无需操作
			continue
		}
		if x < 0 || i+k > n { // 无法操作
			return false
		}
		sumD -= x // 直接加到 sumD 中
		d[i+k] += x
	}
	return true
}
