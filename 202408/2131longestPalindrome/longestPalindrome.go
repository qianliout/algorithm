package main

func main() {

}

func longestPalindrome(words []string) int {
	cnt := make([][]int, 26)
	for i := range cnt {
		cnt[i] = make([]int, 26)
	}
	for _, ch := range words {
		a := int(ch[0]) - int('a')
		b := int(ch[1]) - int('a')
		cnt[a][b]++
	}
	ans := 0
	odd := 0 // // 是否有一个出现奇数次的 AA 类型字符串
	for i := 0; i < 26; i++ {
		c := cnt[i][i] // 两个字母都是相同字符的字符个数
		ans = ans + c - c%2
		odd |= c & 1
		for j := i + 1; j < 26; j++ {
			ans += min(cnt[i][j], cnt[j][i]) * 2
		}
	}
	return (ans + odd) * 2
}

// 对于两个字母相同的情况，与上述类似，我们可以选择偶数个 AA 对称添加到当前回文串左右两侧。如果某个 AA 出现了奇数次，我们还可以将其添加到当前回文串的正中。
