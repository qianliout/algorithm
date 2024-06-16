package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumPushes("avghdcyinjmbopwtrfslzkuxeq"))
}

// 字母不一样的，所以可以直播模拟
func minimumPushes1(word string) int {
	n := len(word)
	ans := 0
	push := 1
	for n > 0 {
		ans += min(n, 8) * push
		n -= 8
		push++
	}
	return ans
}

// 有重复字母且数据量大,只有小写字母
func minimumPushes(word string) int {
	nums := make([]int, 26)
	for _, c := range word {
		nums[int(c-'a')]++
	}

	sort.Ints(nums)
	ans := 0
	push := 1

	i := 25
	for i >= 0 {
		for k := i; k >= max(0, i-8+1); k-- {
			ans += nums[k] * push
		}
		push++
		i -= 8
	}
	return ans
}
