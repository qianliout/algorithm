package main

func main() {

}

func minFlips(a int, b int, c int) int {
	ans := 0
	for a > 0 || b > 0 || c > 0 {
		aa, bb, cc := a&1, b&1, c&1
		if cc == 0 {
			ans += aa + bb
		}
		if cc == 1 && (aa == 0 && bb == 0) {
			ans++
		}
		a, b, c = a>>1, b>>1, c>>1
	}
	return ans
}

/*
给你三个正整数 a、b 和 c。
你可以对 a 和 b 的二进制表示进行位翻转操作，返回能够使按位或运算   a OR b == c  成立的最小翻转次数。
「位翻转操作」是指将一个数的二进制表示任何单个位上的 1 变成 0 或者 0 变成 1 。
*/
