package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(beautifulIndices("isawsquirrelnearmysquirrelhouseohmy", "my", "squirrel", 15)) // 16,33
	fmt.Println(beautifulIndices("frkxslnnn", "rkxsl", "n", 2))                                // []
	fmt.Println(beautifulIndices("lahhnlwx", "hhnlw", "ty", 6))                                // []
}

/*
给你一个下标从 0 开始的字符串 s 、字符串 a 、字符串 b 和一个整数 k 。

如果下标 i 满足以下条件，则认为它是一个 美丽下标：
0 <= i <= s.length - a.length
s[i..(i + a.length - 1)] == a
存在下标 j 使得：
0 <= j <= s.length - b.length
s[j..(j + b.length - 1)] == b
|j - i| <= k
以数组形式按 从小到大排序 返回美丽下标。
*/

func beautifulIndices(s string, a string, b string, k int) []int {
	aa, bb := make([]int, 0), make([]int, 0)
	ss := []byte(s)

	n, na, nb := len(s), len(a), len(b)
	for i := 0; i < len(ss); i++ {
		if i+na <= n && string(ss[i:i+na]) == a {
			aa = append(aa, i)
		}
		if i+nb <= n && string(ss[i:i+nb]) == b {
			bb = append(bb, i)
		}
	}

	ans := make([]int, 0)
	if len(bb) == 0 {
		return ans
	}
	for i := 0; i < len(aa); i++ {
		for j := 0; j < len(bb); j++ {
			if abs(aa[i]-bb[j]) <= k {
				ans = append(ans, aa[i])
				break
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
