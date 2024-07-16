package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSmallestString(3, 27))
}

func getSmallestString(n int, k int) string {
	// 我们先将字符串的每个字符都初始化为 'a'，此时剩余的数值为 d=k−n。
	// 接着从后往前遍历字符串，每次贪心地将当前位置的字符替换为能够使得剩余的数字最小的字符 'z'，直到剩余的数字 d 不超过 25。最后将剩余的数字加到我们遍历到的位置上即可。
	ans := make([]byte, n)
	for i := range ans {
		ans[i] = 'a'
	}
	i, d := n-1, k-n
	for d > 25 {
		ans[i] = 'z'
		d -= 25 // 因为最开始已经加了一个'a'了
		i--
	}
	ans[i] += byte(d)

	return string(ans)
}
