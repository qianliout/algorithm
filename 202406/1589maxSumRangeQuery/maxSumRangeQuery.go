package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maxSumRangeQuery([]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 1}}))
	fmt.Println(maxSumRangeQuery([]int{4, 5, 1}, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}

func maxSumRangeQuery(nums []int, requests [][]int) int {
	n := len(nums)
	d := make([]int, n)
	for _, ch := range requests {
		x, y := ch[0], ch[1]+1
		if x >= 0 && x < n {
			d[x]++
		}
		if y >= 0 && y < n {
			d[y]--
		}
	}
	pre := make([]int, n)
	pre[0] = d[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + d[i]
	}
	// 两个都升序排列，最大的查询对应着最大数，这样就可以获取最大值
	sort.Ints(nums)
	sort.Ints(pre)
	base := int(math.Pow10(9)) + 7
	res := 0
	// 倒序遍历

	// for i := n - 1; i >= 0; i-- {
	// 	// pre[i] 等于0 表示查询已经结束了，因为前面排序了的
	// 	// if pre[i] == 0 {
	// 	// 	break
	// 	// }
	// 	res = (res + nums[i]*pre[i]) % base
	// }
	// 正序遍历，倒序遍历都可以的
	for i := 0; i < n; i++ {
		res = (res + nums[i]*pre[i]) % base
	}
	return res % base
}
