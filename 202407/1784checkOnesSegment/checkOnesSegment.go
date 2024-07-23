package main

func main() {

}

// 感谢评论区的朋友给我的灵感，这个题就是判断字符串是否由左边连续的1和右边连续的0两部分组成，其他的构成都不合法。所以我们只需要找出两个端点然后判断位置就可以了
func checkOnesSegment(s string) bool {
	n := len(s)
	l, r := 0, n-1
	for l < n {
		if s[l] != byte('1') {
			break
		}
		l++
	}
	for r >= 0 {
		if s[r] != byte('0') {
			break
		}
		r--
	}
	return l-r == 1
}
