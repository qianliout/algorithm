package main

import (
	"math"
)

func main() {

}

type pair struct {
	x   int
	cnt int
}

func numberOfSubarrays(nums []int) int64 {
	stark := make([]pair, 0)
	// 递减
	// stark = append(stark, pair{x: math.MaxInt32, cnt: 0})

	ans := len(nums)
	for _, ch := range nums {
		// 维护stark
		// 不加扫描就必须判断为空
		for len(stark) > 0 && ch > stark[len(stark)-1].x {
			stark = stark[:len(stark)-1]
		}
		// 递减栈，那么 ch 是以最小元素的身份加入stark
		if (len(stark) > 0 && ch < stark[len(stark)-1].x) || len(stark) == 0 {
			stark = append(stark, pair{x: ch, cnt: 1})
		} else if len(stark) > 0 && ch == stark[len(stark)-1].x {
			ans += stark[len(stark)-1].cnt
			stark[len(stark)-1].cnt++
		}
	}
	return int64(ans)
}

func numberOfSubarrays1(nums []int) int64 {
	stark := make([]pair, 0)
	// 递减
	// 技巧，加入哨兵
	stark = append(stark, pair{x: math.MaxInt32, cnt: 0})

	ans := len(nums)
	for _, ch := range nums {
		// 维护stark
		// 因为加入了哨兵，可以不用判断 stark 是否为空
		for len(stark) > 0 && ch > stark[len(stark)-1].x {
			stark = stark[:len(stark)-1]
		}
		// 递减栈，那么 ch 是以最小元素的身份加入stark
		if len(stark) > 0 && ch < stark[len(stark)-1].x {
			stark = append(stark, pair{x: ch, cnt: 1})
		} else if len(stark) > 0 && ch == stark[len(stark)-1].x {
			ans += stark[len(stark)-1].cnt
			stark[len(stark)-1].cnt++
		}
	}
	return int64(ans)
}

/*

例如 nums=[4,3,1,2,1]，在从左到右遍历的过程中，由于 2 的出现，左边的 1 永远不可能与右边的 1 组成一个题目要求的子数组。所以当遍历到 2 时，
左边的 1 就是无用数据了，可以清除。清除后我们会得到一个从左到右递减的数据结构。

- 初始化答案等于 n，因为每个元素可以单独组成一个长为 1 的子数组，满足题目要求。
- 维护一个底大顶小的单调栈，记录元素及其出现次数。
- 从左到右遍历 nums
- 只要 x=nums[i] 大于栈顶，就把栈顶出栈。
- 如果 x 小于栈顶，把 x 及其出现次数 1 入栈。
- 如果 x 等于栈顶，设栈顶记录的出现次数为 cnt，那么 x 可以和左边 cnt个 x 组成 cnt 个满足要求的子数组，把答案增加 cnt，然后把 cnt 加一。
*/
