package main

import (
	"fmt"
)

func main() {
	//fmt.Println(repeatLimitedString("cczazcc", 3))
	fmt.Println(repeatLimitedString("aababab", 2))
}

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := make([]int, 26)
	for _, ch := range s {
		idx := int(ch - 'a')
		cnt[idx]++
	}
	ans := make([]byte, 0)
	for {
		find := false
		for i := 25; i >= 0; i-- {
			for cnt[i] > 0 {
				find = true
				mi := min(cnt[i], repeatLimit)
				for j := 0; j < mi; j++ {
					ans = append(ans, byte(i+'a'))
				}
				cnt[i] -= mi
				if cnt[i] > 0 {
					find2 := false
					for j := i - 1; j >= 0; j-- {
						if cnt[j] > 0 {
							ans = append(ans, byte(j+'a'))
							cnt[j]--
							find2 = true
							break
						}
					}
					if !find2 {
						return string(ans)
					}
				}
			}
		}
		if !find {
			break
		}
	}
	return string(ans)
}
