package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

}

/*
把所有数从大到小排列，保证同等长度情况下值最大。
使用一个数组mods， mods[i]mods[i] 表示对3求余结果为i的集合。
再把所有数加起来对3求余:

        如果余数为0将所有数从大到小排列直接返回
        如果余数为1则可以删掉一个余数为1或者两个余数为2的数
        如果余数为2则可以删掉两个余数为1或者一个余数为2的数

为了保证结果最大尽量少删保留较多的位数，且删除最小的
最后将所有数组合在一起从大到小排列并返回.
特殊判断答案为0的情况
*/

func largestMultipleOfThree(digits []int) string {
	sort.Slice(digits, func(i, j int) bool { return digits[i] > digits[j] })
	sum := 0
	mod := make([][]int, 3)
	for _, ch := range digits {
		sum += ch
		mod[ch%3] = append(mod[ch%3], ch)
	}
	for k := range mod {
		sort.Slice(mod[k], func(i, j int) bool { return mod[k][i] > mod[k][j] })
	}

	if sum%3 == 0 {
		return gen(digits)
	}
	m := sum % 3
	if len(mod[m]) > 0 {
		n := len(mod[m])
		mod[m] = mod[m][:n-1]
	} else {
		if len(mod[3-m]) < 2 {
			return ""
		}
		n := len(mod[3-m])
		mod[3-m] = mod[3-m][:n-2]
	}
	ans := make([]int, 0)
	for _, ch := range mod {
		for _, x := range ch {
			ans = append(ans, x)
		}
	}
	sort.Slice(ans, func(i, j int) bool { return ans[i] > ans[j] })
	return gen(ans)
}

func gen(nums []int) string {
	ans := make([]string, 0)
	for _, ch := range nums {
		ans = append(ans, fmt.Sprintf("%d", ch))
	}
	res := strings.Join(ans, "")
	// 应对[0,0,0,0]，这个情况
	if len(nums) > 0 && nums[0] == 0 {

		return "0"
	}

	return res
}
