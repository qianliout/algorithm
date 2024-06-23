package main

import (
	"fmt"
)

func main() {
	fmt.Println(threeEqualParts([]int{1, 0, 1, 0, 1}))
	fmt.Println(threeEqualParts([]int{0, 1, 0, 1, 1}))
	fmt.Println(threeEqualParts([]int{0, 0, 0, 1, 1, 1}))
	fmt.Println(threeEqualParts([]int{0, 0, 0, 0, 0}))
	fmt.Println(threeEqualParts([]int{0, 0, 0, 0, 0}))
}

// 会超时
func threeEqualParts1(nums []int) []int {
	n := len(nums)
	for i := 1; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			a := cal(nums[:i])
			b := cal(nums[i:j])
			c := cal(nums[j:])
			if a == b && b == c {
				return []int{i - 1, j}
			}
		}
	}
	return []int{-1, -1}
}
func cal(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans*2 + nums[i]
	}
	return ans
}

func threeEqualParts(nums []int) []int {
	if len(nums) < 3 {
		return []int{-1, -1}
	}
	cnt := 0
	for _, ch := range nums {
		cnt += ch
	}
	if cnt == 0 {
		return []int{0, len(nums) - 1}
	}
	if cnt%3 != 0 {
		return []int{-1, -1}
	}
	x, y, z := find(nums, 1), find(nums, cnt/3+1), find(nums, cnt/3*2+1)
	for z < len(nums) && nums[x] == nums[y] && nums[y] == nums[z] {
		x++
		y++
		z++
	}
	if z == len(nums) {
		return []int{x - 1, y}
	}

	return []int{-1, -1}
}

func find(nums []int, k int) int {
	ans := 0
	for i, ch := range nums {
		ans += ch
		if ans == k {
			return i
		}
	}
	return 0
}
