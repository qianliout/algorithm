package main

func main() {

}

func largestPalindromic(num string) string {
	cnt := make([]int, 10)
	for _, ch := range num {
		idx := int(ch) - int('0')
		cnt[idx]++
	}
	if cnt[0] == len(num) {
		return "0"
	}
	ans := make([]byte, 0)
	for i := 9; i >= 0; i-- {
		if i == 0 && len(ans) == 0 {
			continue
		}

		for k := 0; k < cnt[i]/2; k++ {
			ans = append(ans, byte('0'+i))
		}
	}
	k := len(ans) - 1
	// 最大回文 整数
	for j := 9; j >= 0; j-- {
		if cnt[j]&1 == 1 {
			ans = append(ans, byte('0'+j))
			break
		}
	}
	for ; k >= 0; k-- {
		ans = append(ans, ans[k])
	}
	return string(ans)
}
