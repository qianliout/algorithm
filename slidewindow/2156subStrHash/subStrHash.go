package main

func main() {

}

// 倒着遍历是这一题的关键
func subStrHash(s string, power int, mod int, k int, hashValue int) string {
	n := len(s)
	hash, pk := 0, 1
	for i := n - 1; i >= n-k; i-- {
		// 将当前字符转换为0到25之间的整数（通过 s[i]&31 实现）
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
		}
	}
	return ans
}
