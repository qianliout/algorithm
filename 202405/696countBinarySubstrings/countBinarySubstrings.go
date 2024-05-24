package main

import "fmt"

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}

func countBinarySubstrings(s string) int {
	ans, preCnt, curCnt := 0, 0, 0
	for i := 0; i < len(s); i++ {
		curCnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			ans += min(preCnt, curCnt)
			preCnt = curCnt
			curCnt = 0
		}
	}
	return ans
}
