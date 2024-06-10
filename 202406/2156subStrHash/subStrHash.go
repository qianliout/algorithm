package main

import (
	"fmt"
)

func main() {
	// fmt.Println(subStrHash2("leetcode", 7, 20, 2, 0))
	// fmt.Println(subStrHash("leetcode", 7, 20, 2, 0))
	// fmt.Println(subStrHash2("fbxzaad", 31, 100, 3, 32))
	// fmt.Println(subStrHash("fbxzaad", 31, 100, 3, 32))
	// fmt.Println(subStrHash2("bzzrtrrpppigevriaooetwawtnfwddgdvoldxucsbyaufhygdxpnxupmvwbryzlgiuierypzqvwiywqlwiygyj", 76, 4, 60, 2))
	fmt.Println(subStrHash("bzzrtrrpppigevriaooetwawtnfwddgdvoldxucsbyaufhygdxpnxupmvwbryzlgiuierypzqvwiywqlwiygyj", 76, 4, 60, 2))
}

// 倒着遍历是这一题的关键
func subStrHash2(s string, power int, mod int, k int, hashValue int) string {
	n := len(s)
	hash, pk := 0, 1
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + int(s[i]&31)) % mod
		pk = pk * power % mod
	}
	ans := ""
	if hash == hashValue {
		// 从后向前，不能直接返回
		ans = s[n-k:]
	}
	for i := n - k - 1; i >= 0; i-- {
		hash = (hash*power + int(s[i]&31) - pk*int(s[i+k]&31)%mod + mod) % mod
		if hash == hashValue {
			// 从后向前，不能直接返回
			ans = s[i : i+k]
			fmt.Println(ans)
		}
	}
	return ans
}

func subStrHash(s string, power, mod, k, hashValue int) (ans string) {
	n := len(s)
	hash, pk := 0, 1
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + int(s[i]&31)) % mod
		pk = pk * power % mod
	}
	if hash == hashValue {
		ans = s[n-k:]
	}
	// 向左滑窗
	for i := n - k - 1; i >= 0; i-- {
		// 计算新的哈希值，注意 +mod 防止计算出负数
		hash = (hash*power + int(s[i]&31) - pk*int(s[i+k]&31)%mod + mod) % mod
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return
}
