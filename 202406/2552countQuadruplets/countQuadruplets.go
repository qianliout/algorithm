package main

import (
	"fmt"
)

func main() {
	fmt.Println(countQuadruplets([]int{1, 3, 2, 4, 5}))
	fmt.Println(countQuadruplets([]int{2, 5, 3, 1, 4}))
}

/*
4 <= nums.length <= 4000
1 <= nums[i] <= nums.length
nums 中所有数字 互不相同 ，nums 是一个排列

这就说明了 nums 中的数据就是1--->n

在 k 右侧的比 nums[j] 大的元素个数，记作 great[k][nums[j]]；
在 j 左侧的比 nums[k]小的元素个数，记作 less[j][nums[k]]。
对于固定的 j 和 k，根据乘法原理，对答案的贡献为

	less[j][nums[k]]⋅great[k][nums[j]]
*/
func countQuadruplets(nums []int) int64 {
	n := len(nums)
	great := make([][]int, n+1)
	less := make([][]int, n+1)
	for i := 0; i < n; i++ {
		great[i] = make([]int, n+1)
		less[i] = make([]int, n+1)
	}
	// 倒着找比 k 大的数的个数
	for k := n - 2; k >= 2; k-- {
		// 上面题目的数据范围决定了可以这样做
		// 比如 [1,2,3,4],当枚举到4的时候，我们就知道，对于1，2，3来说，后面肯定有一个4
		great[k] = append([]int{}, great[k+1]...)
		// 由数据范围决定了，可以这样写
		// 说明，对于小于等于 nums[k+1]-1的数（假如记为 x），就有一个比x大的数了
		for x := 1; x < nums[k+1]; x++ {
			great[k][x]++ // x < nums[k+1]，对于 x，大于它的数的个数 +1
		}
		// 下面这样写也是对的，但是没有上面这写好理解
		// for x := nums[k+1] - 1; x > 0; x-- {
		// 	great[k][x]++ // x < nums[k+1]，对于 x，大于它的数的个数 +1
		// }
	}
	// 同理维护 比j是小的数
	for j := 1; j <= n-2; j++ {
		less[j] = append([]int{}, less[j-1]...)
		for x := n; x > nums[j-1]; x-- {
			less[j][x]++
		}

		// for x := nums[j-1] + 1; x <= n; x++ {
		// 	less[j][x]++
		// }
	}

	ans := 0
	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				ans += great[k][nums[j]] * less[j][nums[k]]
			}
		}
	}
	return int64(ans)
}
