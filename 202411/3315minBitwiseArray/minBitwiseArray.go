package main

import (
	"fmt"
)

func main() {
	fmt.Println(minBitwiseArray([]int{2, 3, 5, 7})) // -1,1,4,3

}

func minBitwiseArray(nums []int) []int {
	// trick的做法，就是直接改nums
	n := len(nums)
	ans := make([]int, n)
	// 根据下面的套路，如是最后一个0后面没有1了，那就不能获取到答案，也就是说偶数是没有答案的
	// 但是本题中 nums[i] 是质数，也只有一个2是偶数
	for i, ch := range nums {
		if ch%2 == 0 {
			ans[i] = -1
			continue
		}

		t := ^ch
		lb := t & (-t)
		a := ch ^ (lb >> 1)
		ans[i] = int(a)
	}
	return ans
}

// ns[i] OR (ans[i] + 1) == nums[i]
// x OR (x + 1) 就是把 x 最右边的一个0变成 1
// 如果 x|(x+1) = 101111,那么x 的值只能是
// 101110,101101,101011,100111
// 题目要求x最小，那结果就只能是 100111,也就是最右边那一个0后面的一个1变成0
