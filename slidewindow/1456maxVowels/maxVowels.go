package main

import (
	"bytes"
)

func main() {

}

func maxVowels2(s string, k int) int {
	wid := 0
	l, r, n := 0, 0, len(s)
	ans := 0
	ae := []byte("aeiou")
	for l <= r && r < n {
		// 入
		idx := bytes.IndexByte(ae, s[r])
		if idx != -1 {
			wid++
		}
		// 更新答案
		if r-l+1 == k {
			ans = max(ans, wid)
		}
		// 出
		for r-l+1 >= k {
			pr := bytes.IndexByte(ae, s[l])
			if pr != -1 {
				wid--
			}
			l++
		}
		r++
	}
	return ans
}

func maxVowels(s string, k int) int {
	ans, cnt := 0, 0
	le, ri, n := 0, 0, len(s)
	for le <= ri && ri < n {
		// 入
		cnt += check(s[ri])
		// 然后马上右移窗口，不然容易忘记
		ri++
		// 更新答案，这里是定长窗口哦
		// 一定要注意，此时的 ri 已经加1了
		if ri-le >= k {
			ans = max(ans, cnt)
		}
		// 出窗口
		for ri-le >= k {
			cnt -= check(s[le])
			le++
		}
	}
	return ans
}

func maxVowels1(s string, k int) int {
	ans, cnt := 0, 0
	le, ri, n := 0, 0, len(s)
	for le <= ri && ri < n {
		// 入
		cnt += check(s[ri])
		// 然后马上右移窗口，不然容易忘记
		ri++
		for ri-le >= k {
			ans = max(ans, cnt)
			cnt -= check(s[le])
			le++
		}
	}
	return ans
}

func check(b byte) int {
	ae := []byte("aeiou")
	id := bytes.IndexByte(ae, b)
	if id != -1 {
		return 1
	}
	return 0
}
