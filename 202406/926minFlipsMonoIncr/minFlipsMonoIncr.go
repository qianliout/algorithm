package main

import (
	"strings"
)

func main() {

}

func minFlipsMonoIncr(s string) int {
	/*
		因此，我们维护两个变量 lb 和 ra 分别表示 s[0,..,i−1] 中字符 bbb 的个数以及 s[i+1,..n−1]中字符 a 的个数，那么我们需要删除的字符数为 lb+ra。枚举过程中，更新变量 lb 和 ra。
	*/
	lb := 0
	ra := strings.Count(s, "0")

	ans := len(s)
	for _, ch := range s {
		if ch == '0' {
			ra--
		}
		ans = min(ans, ra+lb)
		if ch == '1' {
			lb++
		}
	}
	return ans
}
