package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
	fmt.Println(maxUniqueSplit("aba"))
	fmt.Println(maxUniqueSplit("aa"))
	fmt.Println(maxUniqueSplit("wwwzfvedwfvhsww"))
}

func maxUniqueSplit(s string) int {
	ss := []byte(s)
	cnt := make(map[string]bool)
	var dfs func(i int) int
	n := len(s)
	mem := make([]int, n)
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		// 不能这样用缓存，因为值还和 cnt 里的数据有关，如果要用缓存，得把 cnt 也缓存了
		// if mem[i] != -1 {
		// 	return mem[i]
		// }
		res := 0

		for j := i + 1; j <= n; j++ {
			key := string(ss[i:j])
			if cnt[key] {
				continue
			}
			cnt[key] = true
			res = max(res, dfs(j)+1)
			cnt[key] = false
		}
		// mem[i] = res
		return res
	}
	res := dfs(0)
	return res
}
