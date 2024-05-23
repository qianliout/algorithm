package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumScore([]int{1, 4, 3, 7, 4, 5}, 3))
	fmt.Println(maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0))
	fmt.Println(maximumScore([]int{8407, 5626, 9236, 4362, 9678, 6568, 4170, 5691, 7865, 4067, 2094,
		9451, 9667, 1400, 5093, 6191, 7286, 7368, 6461, 4309, 9751, 3672, 4165, 4940, 3726, 7003, 6263,
		4250, 1950, 9536, 61, 1486, 6009, 6977, 7084, 2472, 7921, 1913, 117, 3543, 5075, 1582,
		7987, 6710, 1331, 3023, 6856, 1125, 1475, 4158, 3422, 7864, 9178, 7255, 4997, 2128, 5441,
		5910, 6719, 1308, 4444, 9746}, 30)) // 3782
}

func maximumScore(nums []int, k int) int {
	nums = append([]int{math.MinInt32}, nums...)
	nums = append(nums, math.MinInt32)

	// 单调递增
	stark := make([]int, 0)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for len(stark) > 0 && nums[stark[len(stark)-1]] > nums[i] {
			top := nums[stark[len(stark)-1]]
			stark = stark[:len(stark)-1]
			left := stark[len(stark)-1] // 因为在开始时增加了哨兵节点，所以不会越界
			right := i - 1
			// 这里的判断，真的有点麻烦啊
			// [left   k    right]
			// 此时 left是所选左边界的左边一个index,不是所选的区间的左边 index,所以 left+1<=k+1
			// 右边 right 就是所选的区间的右边的 Index
			if left+1 <= k+1 && k+1 <= right {
				ans = max(ans, top*(right-left))
			}
		}
		stark = append(stark, i)
	}

	return ans
}
