package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(palindromePartition("abc", 2))
	fmt.Println(palindromePartition("aabbc", 3))
	fmt.Println(palindromePartition("leetcode", 8))
	fmt.Println(palindromePartition("fyhowoxzyrincxivwarjuwxrwealesxsimsepjdqsstfggjnjhilvrwwytbgsqbpnwjaojfnmiqiqnyzijfmvekgakefjaxryyml", 32))
}
func palindromePartition(s string, k int) int {
	return dfs([]byte(s), k)
}

func dfs(ss []byte, k int) int {
	if k <= 0 {
		return 0
	}
	if len(ss) <= k {
		return 0
	}
	if k == 1 {
		return change(ss)
	}
	res := math.MaxInt
	for i := 1; i <= len(ss); i++ {
		res = min(res, change(ss[:i])+dfs(ss[i:], k-1))
	}
	return res
}

func change(ss []byte) int {
	le, ri := 0, len(ss)-1
	ans := 0
	for le < ri {
		if ss[le] != ss[ri] {
			ans++
		}
		le++
		ri--
	}
	return ans
}
