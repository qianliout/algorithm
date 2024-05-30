package main

func main() {

}

func reorganizeString(s string) string {
	count := make([]int, 26)
	// mac 最多的字符是有多少个
	// mab 最多的字符是那一个
	mac, mab := 0, byte(0)
	for _, ch := range s {
		count[ch-'a']++
		if count[ch-'a'] > mac {
			mac = count[ch-'a']
			mab = byte(ch)
		}
	}
	// 这种情况下就不得行
	// n:=len(s)
	// 如果 n 是奇数，那么 mac 最大是：n/2+1 (最中间放最多的元素)
	// 如果 n 是偶数，那么 mac 最大是 n/2
	if len(s)%2 == 1 && mac > len(s)/2+1 {
		return ""
	}
	if len(s)%2 == 0 && mac > len(s)/2 {
		return ""
	}
	// 这样的判断更简洁
	// if mac > len(s)-mac+1 {
	// 	return ""
	// }
	ans := make([]byte, len(s))
	i := 0
	for mac > 0 {
		ans[i] = byte(mab)
		i += 2
		mac--
	}
	count[mab-'a'] = 0

	for b, c := range count {
		for c > 0 {
			// 这里要理解，最多的那个字符不一定会填充到最后,所以其他的字符会接着填充，只有当填充到最后时才回到最前面
			if i >= len(s) {
				i = 1 // 最多的值是从0开始填的，其他从1开始填
			}
			ans[i] = byte('a' + b)
			i += 2
			c--
		}
	}
	return string(ans)
}
