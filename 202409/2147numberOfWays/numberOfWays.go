package main

func main() {

}

func numberOfWays(corridor string) int {
	ans := 1
	cntS := 0
	pre := 0
	mod := int(1e9 + 7)
	for i, ch := range corridor {
		if ch == 'S' {
			cntS++
			if cntS >= 3 && cntS%2 != 0 {
				ans = ans * (i - pre)
				ans = ans % mod
			}
			pre = i
		}
	}
	// 这步判断容易出错
	// 最后还必须有2个位置
	if cntS > 0 && cntS%2 == 0 {
		return ans
	}
	return 0
}
